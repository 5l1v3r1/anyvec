package cuda

import (
	"runtime"
	"unsafe"

	"github.com/unixpickle/anyvec"
)

// A Creator32 implements anyvec.Creator for float32
// numerics.
type Creator32 struct {
	handle *Handle
}

// NewCreator32 creates a Creator32 that uses a given CUDA
// handle.
func NewCreator32(h *Handle) *Creator32 {
	return &Creator32{handle: h}
}

// MakeNumeric creates a float32.
func (c *Creator32) MakeNumeric(x float64) anyvec.Numeric {
	return float32(x)
}

// MakeNumericList creates a []float32.
func (c *Creator32) MakeNumericList(x []float64) anyvec.NumericList {
	res := make([]float32, len(x))
	for i, k := range x {
		res[i] = float32(k)
	}
	return res
}

// MakeVector creates a zero'd out anyvec.Vector.
func (c *Creator32) MakeVector(size int) anyvec.Vector {
	buf, err := newBuffer(c.handle, size*4)
	if err != nil {
		panic(err)
	}
	if err := buf.Clear(); err != nil {
		panic(err)
	}
	return &vector32{
		handle: c.handle,
		buffer: buf,
	}
}

// MakeVectorData creates an anyvec.Vector with the
// specified contents.
func (c *Creator32) MakeVectorData(dObj anyvec.NumericList) anyvec.Vector {
	d := dObj.([]float32)
	buf, err := newBuffer(c.handle, len(d)*4)
	if err != nil {
		panic(err)
	}
	if err := buf.Set32(d); err != nil {
		panic(err)
	}
	return &vector32{
		handle: c.handle,
		buffer: buf,
	}
}

// Concat concatenates vectors.
func (c *Creator32) Concat(v ...anyvec.Vector) anyvec.Vector {
	bufs := make([]*buffer, len(v))
	for i, x := range v {
		bufs[i] = x.(*vector32).buffer
	}
	buf, err := newBufferConcat(c.handle, bufs)
	if err != nil {
		panic(err)
	}
	return &vector32{
		handle: c.handle,
		buffer: buf,
	}
}

type vector32 struct {
	handle *Handle
	buffer *buffer
}

func (v *vector32) Len() int {
	return v.buffer.Len() / 4
}

func (v *vector32) Data() anyvec.NumericList {
	res := make([]float32, v.Len())
	if err := v.buffer.Get32(res); err != nil {
		panic(err)
	}
	return res
}

func (v *vector32) SetData(d anyvec.NumericList) {
	if err := v.buffer.Set32(d.([]float32)); err != nil {
		panic(err)
	}
}

func (v *vector32) Copy() anyvec.Vector {
	newBuff, err := newBuffer(v.handle, v.buffer.size)
	if err != nil {
		panic(err)
	}
	if err := newBuff.Set(v.buffer); err != nil {
		panic(err)
	}
	return &vector32{
		handle: v.handle,
		buffer: newBuff,
	}
}

func (v *vector32) Slice(start, end int) anyvec.Vector {
	if start < 0 || end < 0 {
		panic("indices must be non-negative")
	}
	if end < start {
		panic("invalid range: end < start")
	}
	if end > v.Len() {
		panic("end out of bounds")
	}
	buf, err := newBuffer(v.handle, (end-start)*4)
	if err != nil {
		panic(err)
	}
	buf.Set(&buffer{
		size: (end - start) * 4,
		ptr:  unsafe.Pointer(uintptr(v.buffer.ptr) + uintptr(4*start)),
	})
	runtime.KeepAlive(v.buffer)
	return &vector32{
		handle: v.handle,
		buffer: buf,
	}
}

func (v *vector32) Scale(s anyvec.Numeric) {
	v.handle.sscal(v.Len(), s.(float32), v.buffer.ptr)
}

func (v *vector32) AddScaler(s anyvec.Numeric) {
	constVec, err := newBuffer(v.handle, v.buffer.size)
	if err != nil {
		panic(err)
	}
	if err := constVec.SetRepeated32(s.(float32)); err != nil {
		panic(err)
	}
	v.Add(&vector32{handle: v.handle, buffer: constVec})
}

func (v *vector32) Dot(v1 anyvec.Vector) anyvec.Numeric {
	v.assertMatch(v1)
	return v.handle.sdot(v.Len(), v.buffer.ptr, v1.(*vector32).buffer.ptr)
}

func (v *vector32) Add(v1 anyvec.Vector) {
	v.assertMatch(v1)
	v.handle.saxpy(v.Len(), 1, v1.(*vector32).buffer.ptr, v.buffer.ptr)
}

func (v *vector32) Sub(v1 anyvec.Vector) {
	v.assertMatch(v1)
	v.handle.saxpy(v.Len(), -1, v1.(*vector32).buffer.ptr, v.buffer.ptr)
}

func (v *vector32) Mul(v1 anyvec.Vector) {
	v.assertMatch(v1)
	v.handle.mul(v.Len(), v.buffer.ptr, v1.(*vector32).buffer.ptr)
}

func (v *vector32) Div(v1 anyvec.Vector) {
	v.assertMatch(v1)
	v.handle.div(v.Len(), v.buffer.ptr, v1.(*vector32).buffer.ptr)
}

func (v *vector32) Gemm(transA, transB bool, m, n, k int, alpha anyvec.Numeric, a anyvec.Vector,
	lda int, b anyvec.Vector, ldb int, beta anyvec.Numeric, ldc int) {
	validateGemm(transA, transB, m, n, k, a.Len(), lda, b.Len(), ldb, v.Len(), ldc)
	v.handle.sgemm(transA, transB, m, n, k, alpha.(float32), a.(*vector32).buffer.ptr,
		lda, b.(*vector32).buffer.ptr, ldb, beta.(float32), v.buffer.ptr, ldc)
}

func (v *vector32) Exp() {
	v.handle.exp(v.Len(), v.buffer.ptr)
}

func (v *vector32) Tanh() {
	v.handle.tanh(v.Len(), v.buffer.ptr)
}

func (v *vector32) Sin() {
	v.handle.sin(v.Len(), v.buffer.ptr)
}

func (v *vector32) ClipPos() {
	v.handle.clipPos(v.Len(), v.buffer.ptr)
}

func (v *vector32) assertMatch(v1 anyvec.Vector) {
	if v.Len() != v1.Len() {
		panic("sizes do no match")
	}
}
