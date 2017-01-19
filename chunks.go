package anyvec

import "fmt"

// A ChunkAdder can add a different scaler to each
// contiguous chunk of itself.
// It is equivalent to the pseudo-code:
//
//     chunkSize := v.Len() / scalers.Len()
//     v[i] += scalers[i / chunkSize];
//
// The length of the receiver must be divisible by the
// length of the scaler vector, since each chunk must be
// the same length.
type ChunkAdder interface {
	AddChunks(scalers Vector)
}

// AddChunks adds scalers to contiguous chunks of a
// vector.
// If the vector does not implement ChunkAdder, a default
// implementation is used which supports float32 and
// float64 values.
// v and scalers must not be equal.
func AddChunks(v, scalers Vector) {
	if c, ok := v.(ChunkAdder); ok {
		c.AddChunks(scalers)
	} else {
		if v.Len()%scalers.Len() != 0 {
			panic("scaler count must divide vector length")
		}
		data := v.Data()
		switch data := data.(type) {
		case []float32:
			applyAddChunks32(data, scalers.Data().([]float32))
		case []float64:
			applyAddChunks64(data, scalers.Data().([]float64))
		default:
			panic(fmt.Sprintf("unsupported type: %T", data))
		}
		v.SetData(data)
	}
}

func applyAddChunks32(data, scalers []float32) {
	chunkSize := len(data) / len(scalers)
	for i := range data {
		data[i] += scalers[i/chunkSize]
	}
}

func applyAddChunks64(data, scalers []float64) {
	chunkSize := len(data) / len(scalers)
	for i := range data {
		data[i] += scalers[i/chunkSize]
	}
}

// A ChunkScaler can scale contiguous chunks of itself,
// each by differenent scalers.
// It is equivalent to the pseudo-code:
//
//     chunkSize := v.Len() / scalers.Len()
//     v[i] *= scalers[i / chunkSize];
//
// The length of the receiver must be divisible by the
// length of the scaler vector, since each chunk must be
// the same length.
type ChunkScaler interface {
	ScaleChunks(scalers Vector)
}

// ScaleChunks scales contiguous chunks of a vector.
// If the vector does not implement ChunkScaler, a default
// implementation is used which supports float32 and
// float64 values.
// v and scalers must not be equal.
func ScaleChunks(v, scalers Vector) {
	if c, ok := v.(ChunkScaler); ok {
		c.ScaleChunks(scalers)
	} else {
		if v.Len()%scalers.Len() != 0 {
			panic("scaler count must divide vector length")
		}
		data := v.Data()
		switch data := data.(type) {
		case []float32:
			applyScaleChunks32(data, scalers.Data().([]float32))
		case []float64:
			applyScaleChunks64(data, scalers.Data().([]float64))
		default:
			panic(fmt.Sprintf("unsupported type: %T", data))
		}
		v.SetData(data)
	}
}

func applyScaleChunks32(data, scalers []float32) {
	chunkSize := len(data) / len(scalers)
	for i := range data {
		data[i] *= scalers[i/chunkSize]
	}
}

func applyScaleChunks64(data, scalers []float64) {
	chunkSize := len(data) / len(scalers)
	for i := range data {
		data[i] *= scalers[i/chunkSize]
	}
}

// A RepeatedAdder can add to itself the repeated contents
// of another vector.
// It is equivalent to the following pseudo-code:
//
//     v[i] += scalers[i%scalers.Len()]
//
// The length of the scalers will never be zero.
type RepeatedAdder interface {
	AddRepeated(scalers Vector)
}

// AddRepeated adds the repeated form of scalers to v.
// If the vector does not implement RepeatedAdder, a
// default implementation is used.
// v and scalers must not be equal.
func AddRepeated(v, scalers Vector) {
	if r, ok := v.(RepeatedAdder); ok {
		r.AddRepeated(scalers)
	} else {
		if scalers.Len() == 0 {
			panic("repeated vector cannot be empty")
		}
		var joinMe []Vector
		var joinLen int
		for joinLen < v.Len() {
			joinLen += scalers.Len()
			joinMe = append(joinMe, scalers)
		}
		v.Add(v.Creator().Concat(joinMe...).Slice(0, v.Len()))
	}
}

// A RepeatedScaler can scale to itself the repeated
// contents of another vector.
// It is equivalent to the following pseudo-code:
//
//     v[i] *= scalers[i%scalers.Len()]
//
// The length of the scalers will never be zero.
type RepeatedScaler interface {
	ScaleRepeated(scalers Vector)
}

// ScaleRepeated scales the repeated form of scalers to v.
// If the vector does not implement RepeatedScaler, a
// default implementation is used.
// v and scalers must not be equal.
func ScaleRepeated(v, scalers Vector) {
	if r, ok := v.(RepeatedScaler); ok {
		r.ScaleRepeated(scalers)
	} else {
		if scalers.Len() == 0 {
			panic("repeated vector cannot be empty")
		}
		var joinMe []Vector
		var joinLen int
		for joinLen < v.Len() {
			joinLen += scalers.Len()
			joinMe = append(joinMe, scalers)
		}
		v.Mul(v.Creator().Concat(joinMe...).Slice(0, v.Len()))
	}
}

// A RowSummer can sum the rows of a row-major matrix.
// The number of columns must divide the size of the
// vector.
type RowSummer interface {
	SumRows(cols int) Vector
}

// SumRows sums the rows of a row-major matrix.
// If the vector does not implement RowSummer, a default
// implementation is used.
func SumRows(v Vector, cols int) Vector {
	if r, ok := v.(RowSummer); ok {
		return r.SumRows(cols)
	} else {
		if v.Len()%cols != 0 {
			panic("number of columns must divide vector size")
		}
		rows := v.Len() / cols
		oneVec := v.Creator().MakeVector(rows)
		oneVec.AddScaler(v.Creator().MakeNumeric(1))
		oneMat := &Matrix{
			Data: oneVec,
			Rows: rows,
			Cols: 1,
		}
		vMat := &Matrix{
			Data: v,
			Rows: rows,
			Cols: cols,
		}
		outMat := &Matrix{
			Data: v.Creator().MakeVector(cols),
			Rows: cols,
			Cols: 1,
		}
		outMat.Product(true, false, v.Creator().MakeNumeric(1), vMat, oneMat,
			v.Creator().MakeNumeric(0))
		return outMat.Data
	}
}

// A ColSummer can sum the columns of a row-major matrix.
// The number of rows must divide the size of the vector.
type ColSummer interface {
	SumCols(rows int) Vector
}

// SumCols sums the columns of a row-major matrix.
// If the vector does not implement ColSummer, a default
// implementation is used.
func SumCols(v Vector, rows int) Vector {
	if c, ok := v.(ColSummer); ok {
		return c.SumCols(rows)
	} else {
		if v.Len()%rows != 0 {
			panic("number of rows must divide vector size")
		}
		cols := v.Len() / rows
		oneVec := v.Creator().MakeVector(cols)
		oneVec.AddScaler(v.Creator().MakeNumeric(1))
		oneMat := &Matrix{
			Data: oneVec,
			Rows: cols,
			Cols: 1,
		}
		vMat := &Matrix{
			Data: v,
			Rows: rows,
			Cols: cols,
		}
		outMat := &Matrix{
			Data: v.Creator().MakeVector(rows),
			Rows: rows,
			Cols: 1,
		}
		outMat.Product(false, false, v.Creator().MakeNumeric(1), vMat, oneMat,
			v.Creator().MakeNumeric(0))
		return outMat.Data
	}
}
