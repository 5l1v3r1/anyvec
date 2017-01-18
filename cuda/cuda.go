//+build !nocuda

package cuda

/*
#include "cuda.h"
#include "cuda_runtime_api.h"
#include "cublas_v2.h"

cublasOperation_t noTranspose = CUBLAS_OP_N;
cublasOperation_t transpose = CUBLAS_OP_T;
cublasSideMode_t sideModeRight = CUBLAS_SIDE_RIGHT;
cublasSideMode_t sideModeLeft = CUBLAS_SIDE_LEFT;
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

// A Handle manages an internal CUDA context.
type Handle struct {
	loop    *cudaLoop
	kernels *mathKernels
	rand    *randomizer
}

// NewHandle attempts to get a new Handle.
func NewHandle() (*Handle, error) {
	err := createMainLoop()
	if err != nil {
		return nil, err
	}
	res := &Handle{loop: getMainLoop()}
	runtime.SetFinalizer(res, func(obj *Handle) {
		obj.loop.Run(func() {
			if obj.kernels != nil {
				obj.kernels.Destroy()
			}
			if obj.rand != nil {
				obj.rand.Destroy()
			}
		})
	})
	return res, nil
}

func (h *Handle) runWithKernels(f func() error) {
	h.loop.Run(func() {
		var err error
		if h.kernels == nil {
			h.kernels, err = newMathKernels()
			if err != nil {
				panic(err)
			}
		}
		err = f()
		if err != nil {
			panic(err)
		}
	})
}

func (h *Handle) runWithRand(f func() error) {
	h.loop.Run(func() {
		var err error
		if h.kernels == nil {
			h.kernels, err = newMathKernels()
			if err != nil {
				panic(err)
			}
		}
		if h.rand == nil {
			h.rand, err = newRandomizer()
			if err != nil {
				panic(err)
			}
		}
		err = f()
		if err != nil {
			panic(err)
		}
	})
}

// A buffer is an on-device memory buffer.
type buffer struct {
	handle *Handle
	size   int
	ptr    unsafe.Pointer
}

// newBuffer allocates a buffer.
func newBuffer(h *Handle, size int) (res *buffer, err error) {
	h.loop.Run(func() {
		var buf unsafe.Pointer
		err = cudaError("cudaMalloc", C.cudaMalloc(&buf, C.size_t(size)))
		if err == nil {
			res = &buffer{
				handle: h,
				size:   size,
				ptr:    buf,
			}
		}
	})
	if err != nil {
		return nil, err
	}
	runtime.SetFinalizer(res, func(b *buffer) {
		b.handle.loop.Run(func() {
			C.cudaFree(b.ptr)
		})
	})
	return res, nil
}

// newBufferConcat concatenates buffers.
func newBufferConcat(h *Handle, bufs []*buffer) (*buffer, error) {
	var size int
	for _, x := range bufs {
		size += x.size
	}
	buf, err := newBuffer(h, size)
	if err != nil {
		return nil, err
	}
	h.loop.Run(func() {
		var idx uintptr
		for _, x := range bufs {
			dest := unsafe.Pointer(uintptr(buf.ptr) + idx)
			idx += uintptr(x.size)
			err = cudaError("cudaMemcpy", C.cudaMemcpy(dest, x.ptr, C.size_t(x.size),
				C.cudaMemcpyDeviceToDevice))
			if err != nil {
				return
			}
		}
	})
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// Len returns the buffer's length in bytes.
func (b *buffer) Len() int {
	return b.size
}

// Clear zeroes the buffer.
func (b *buffer) Clear() error {
	var err error
	b.handle.loop.Run(func() {
		err = cudaError("cudaMemset", C.cudaMemset(b.ptr, 0, C.size_t(b.size)))
	})
	runtime.KeepAlive(b)
	return err
}

// Set copies the contents of a buffer into b.
func (b *buffer) Set(b1 *buffer) error {
	if b1.size != b.size {
		return errors.New("buffer sizes do not match")
	}
	var res error
	b.handle.loop.Run(func() {
		res = cudaError("cudaMemcpy", C.cudaMemcpy(b.ptr, b1.ptr, C.size_t(b.size),
			C.cudaMemcpyDeviceToDevice))
	})
	runtime.KeepAlive(b1)
	runtime.KeepAlive(b)
	return res
}

// Set32 copies 32-bit floats into the buffer.
func (b *buffer) Set32(src []float32) error {
	res := b.hostToDevice(len(src)*4, unsafe.Pointer(&src[0]))
	runtime.KeepAlive(src)
	return res
}

// SetRepeated32 copies the same 32-bits again and again
// to fill the buffer.
func (b *buffer) SetRepeated32(v float32) error {
	if b.size%4 != 0 {
		panic("size not divisible by 4")
	}
	buf := make([]float32, b.size/4)
	for i := range buf {
		buf[i] = v
	}
	return b.Set32(buf)
}

// Set64 copies 64-bit floats into the buffer.
func (b *buffer) Set64(src []float64) error {
	res := b.hostToDevice(len(src)*8, unsafe.Pointer(&src[0]))
	runtime.KeepAlive(src)
	return res
}

// SetRepeated64 copies the same 64-bits again and again
// to fill the buffer.
func (b *buffer) SetRepeated64(v float64) error {
	if b.size%8 != 0 {
		panic("size not divisible by 8")
	}
	buf := make([]float64, b.size/8)
	for i := range buf {
		buf[i] = v
	}
	return b.Set64(buf)
}

// Get32 copies 32-bit floats out of the buffer.
func (b *buffer) Get32(dst []float32) error {
	res := b.deviceToHost(len(dst)*4, unsafe.Pointer(&dst[0]))
	runtime.KeepAlive(dst)
	return res
}

// Get64 copies 64-bit floats out of the buffer.
func (b *buffer) Get64(dst []float64) error {
	res := b.deviceToHost(len(dst)*8, unsafe.Pointer(&dst[0]))
	runtime.KeepAlive(dst)
	return res
}

func (b *buffer) hostToDevice(size int, src unsafe.Pointer) error {
	if size > b.size {
		panic("buffer overflow")
	}
	var res error
	b.handle.loop.Run(func() {
		res = cudaError("cudaMemcpy", C.cudaMemcpy(b.ptr, src, C.size_t(size),
			C.cudaMemcpyHostToDevice))
	})
	runtime.KeepAlive(b)
	return res
}

func (b *buffer) deviceToHost(size int, dst unsafe.Pointer) error {
	if size > b.size {
		panic("buffer overflow")
	}
	var res error
	b.handle.loop.Run(func() {
		res = cudaError("cudaMemcpy", C.cudaMemcpy(dst, b.ptr, C.size_t(size),
			C.cudaMemcpyDeviceToHost))
	})
	runtime.KeepAlive(b)
	return res
}

func blasTransposeOp(trans bool) C.cublasOperation_t {
	if trans {
		return C.transpose
	}
	return C.noTranspose
}
