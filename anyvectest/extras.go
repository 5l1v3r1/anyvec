package anyvectest

import (
	"math"
	"math/rand"
	"testing"

	"github.com/unixpickle/anyvec"
	"github.com/unixpickle/approb"
)

// TestExtras runs tests for all of the possible extra
// methods the Vector may implement.
func (t *Tester) TestExtras(test *testing.T) {
	test.Run("Exp", t.TestExp)
	test.Run("Sin", t.TestSin)
	test.Run("Tanh", t.TestTanh)
	test.Run("ClipPos", t.TestClipPos)
	test.Run("Max", t.TestMax)
	test.Run("Sum", t.TestSum)
	test.Run("AbsMax", t.TestAbsMax)
	test.Run("AbsSum", t.TestAbsSum)
	test.Run("Norm", t.TestNorm)
	test.Run("LogSoftmax", t.TestLogSoftmax)
	test.Run("AddChunks", t.TestAddChunks)
	test.Run("ScaleChunks", t.TestScaleChunks)
	test.Run("Rand", t.TestRand)
	test.Run("AddRepeated", t.TestAddRepeated)
	test.Run("ScaleRepeated", t.TestScaleRepeated)
	test.Run("Comparisons", t.TestComparisons)
	test.Run("TestComplement", t.TestComplement)
}

// TestExp tests exponentiation.
func (t *Tester) TestExp(test *testing.T) {
	t.testOp(test, func(x float64) float64 {
		return math.Exp(x)
	}, anyvec.Exp)
}

// TestSin tests sine.
func (t *Tester) TestSin(test *testing.T) {
	t.testOp(test, func(x float64) float64 {
		return math.Sin(x)
	}, anyvec.Sin)
}

// TestTanh tests tanh.
func (t *Tester) TestTanh(test *testing.T) {
	t.testOp(test, func(x float64) float64 {
		return math.Tanh(x)
	}, anyvec.Tanh)
}

// TestClipPos tests positive clipping.
func (t *Tester) TestClipPos(test *testing.T) {
	t.testOp(test, func(x float64) float64 {
		return math.Max(0, x)
	}, anyvec.ClipPos)
}

// TestSum tests summation.
func (t *Tester) TestSum(test *testing.T) {
	t.testAgg(test, func(x []float64) float64 {
		var sum float64
		for _, k := range x {
			sum += k
		}
		return sum
	}, anyvec.Sum)
}

// TestMax tests max computation.
func (t *Tester) TestMax(test *testing.T) {
	t.testAgg(test, func(x []float64) float64 {
		max := math.Inf(-1)
		for _, k := range x {
			if k > max {
				max = k
			}
		}
		return max
	}, anyvec.Max)
}

// TestAbsSum tests absolute summation.
func (t *Tester) TestAbsSum(test *testing.T) {
	t.testAgg(test, func(x []float64) float64 {
		var sum float64
		for _, k := range x {
			if k > 0 {
				sum += k
			} else {
				sum -= k
			}
		}
		return sum
	}, anyvec.AbsSum)
}

// TestNorm tests the Euclidean norm.
func (t *Tester) TestNorm(test *testing.T) {
	t.testAgg(test, func(x []float64) float64 {
		var sum float64
		for _, k := range x {
			sum += k * k
		}
		return math.Sqrt(sum)
	}, anyvec.Norm)
}

// TestAbsMax tests absolute max computation.
func (t *Tester) TestAbsMax(test *testing.T) {
	t.testAgg(test, func(x []float64) float64 {
		var max float64
		for _, k := range x {
			x := k
			if x < 0 {
				x = -x
			}
			if x > max {
				max = x
			}
		}
		return max
	}, anyvec.AbsMax)
}

// TestLogSoftmax tests log-domain softmaxing.
func (t *Tester) TestLogSoftmax(test *testing.T) {
	inData := []float64{0.735181, 0.638349, 0.672189, 0.025860, 0.492178,
		0.175069, 0.487960, 0.598212}
	expected := []float64{-1.84798499753082, -1.94481718265820, -1.91097680228496,
		-2.55730580715485, -2.09098781021549, -2.40809688666664, -2.09520595321559,
		-1.98495420701728}
	v := t.vec(inData)
	anyvec.LogSoftmax(v, 0)
	t.assertClose(test, v.Data(), expected)

	expected = []float64{-1.20556282283073, -1.30239482283073, -1.26855482283073,
		-1.91488382283073, -1.34440459198132, -1.66151359198132, -1.34862259198132,
		-1.23837059198132}
	v = t.vec(inData)
	anyvec.LogSoftmax(v, 4)
	t.assertClose(test, v.Data(), expected)

	inData = []float64{0.86568, -0.61489, 0.59658, -0.18710, -1.46797, -0.73938}
	expected = []float64{-0.97016, -2.45073, -1.23925, -2.02294, -3.30380, -2.57522}
	v = t.vec(inData)
	anyvec.LogSoftmax(v, 6)
	t.assertClose(test, v.Data(), expected)
}

// TestAddChunks tests chunk addition.
func (t *Tester) TestAddChunks(test *testing.T) {
	data := []float64{-0.66886, -2.12313, -0.49031, -1.00323, -0.82617, 1.16384, -1.84009,
		-0.24389, 0.25384, -0.78570}
	scalers := []float64{-0.25492, -0.50632}
	expected := make([]float64, len(data))
	for i, x := range data {
		expected[i] = x + scalers[i/5]
	}
	v := t.vec(data)
	anyvec.AddChunks(v, t.vec(scalers))
	t.assertClose(test, v.Data(), expected)
}

// TestScaleChunks tests chunk scaling.
func (t *Tester) TestScaleChunks(test *testing.T) {
	data := []float64{-0.66886, -2.12313, -0.49031, -1.00323, -0.82617, 1.16384, -1.84009,
		-0.24389, 0.25384, -0.78570}
	scales := []float64{-0.25492, -0.50632}
	expected := []float64{0.17051, 0.54123, 0.12499, 0.25574, 0.21061, -0.58927, 0.93167,
		0.12349, -0.12852, 0.39781}
	v := t.vec(data)
	anyvec.ScaleChunks(v, t.vec(scales))
	t.assertClose(test, v.Data(), expected)
}

// TestRand tests random sampling.
func (t *Tester) TestRand(test *testing.T) {
	if testing.Short() {
		test.Skip("not running in short mode")
	}

	const numSamples = 20000
	const vecSize = 15
	r := rand.New(rand.NewSource(1337))

	vec := t.Creator.MakeVector(vecSize)

	names := []string{"Uniform", "Bernoulli", "Normal"}
	dists := []anyvec.ProbDist{anyvec.Uniform, anyvec.Bernoulli, anyvec.Normal}
	samplers := []func() float64{rand.Float64, func() float64 {
		return float64(rand.Intn(2))
	}, rand.NormFloat64}
	for i, name := range names {
		actualSampler := func() float64 {
			if rand.Intn(2) == 0 {
				anyvec.Rand(vec, dists[i], r)
			} else {
				anyvec.Rand(vec, dists[i], nil)
			}
			return t.unlist(vec.Data())[rand.Intn(vecSize)]
		}
		corr := approb.Correlation(numSamples, 0.1, actualSampler, samplers[i])
		if corr < 0.95 {
			test.Errorf("distribution %s had correlation %f", name, corr)
		}
	}
}

// TestAddRepeated tests adding a repeated vector.
func (t *Tester) TestAddRepeated(test *testing.T) {
	v := t.vec([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	v1 := t.vec([]float64{-3, -4, -5})

	anyvec.AddRepeated(v, v1)
	expected := []float64{1 - 3, 2 - 4, 3 - 5, 4 - 3, 5 - 4, 6 - 5, 7 - 3,
		8 - 4, 9 - 5, 10 - 3}
	t.assertClose(test, v.Data(), expected)

	anyvec.AddRepeated(v1, v)
	expected = []float64{-3 + 1 - 3, -4 + 2 - 4, -5 + 3 - 5}
	t.assertClose(test, v1.Data(), expected)

	v1 = t.vec([]float64{-1, -2})
	v = t.vec([]float64{1, 2, 3, 4, 5, 6, 7, 8})
	anyvec.AddRepeated(v, v1)
	expected = []float64{0, 0, 3 - 1, 4 - 2, 5 - 1, 6 - 2, 7 - 1, 8 - 2}
	t.assertClose(test, v.Data(), expected)
}

// TestScaleRepeated tests scaling a repeated vector.
func (t *Tester) TestScaleRepeated(test *testing.T) {
	v := t.vec([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	v1 := t.vec([]float64{-3, -4, -5})

	anyvec.ScaleRepeated(v, v1)
	expected := []float64{1 * -3, 2 * -4, 3 * -5, 4 * -3, 5 * -4, 6 * -5, 7 * -3,
		8 * -4, 9 * -5, 10 * -3}
	t.assertClose(test, v.Data(), expected)

	anyvec.ScaleRepeated(v1, v)
	expected = []float64{-3 * expected[0], -4 * expected[1], -5 * expected[2]}
	t.assertClose(test, v1.Data(), expected)

	v1 = t.vec([]float64{-1, -2})
	v = t.vec([]float64{1, 2, 3, 4, 5, 6, 7, 8})
	anyvec.ScaleRepeated(v, v1)
	expected = []float64{-1, -4, -3, -8, -5, -12, -7, -16}
	t.assertClose(test, v.Data(), expected)
}

// TestComparisons tests the Comparer interface.
func (t *Tester) TestComparisons(test *testing.T) {
	data := []float64{1, -1, 2, -2, 3, -3, 1.5, 7, 0, 1}
	v := t.vec(data)
	anyvec.GreaterThan(v, t.num(1.5))
	expected := []float64{0, 0, 1, 0, 1, 0, 0, 1, 0, 0}
	t.assertClose(test, v.Data(), expected)

	v = t.vec(data)
	anyvec.LessThan(v, t.num(-1))
	expected = []float64{0, 0, 0, 1, 0, 1, 0, 0, 0, 0}
	t.assertClose(test, v.Data(), expected)

	v = t.vec(data)
	anyvec.EqualTo(v, t.num(2))
	expected = []float64{0, 0, 1, 0, 0, 0, 0, 0, 0, 0}
	t.assertClose(test, v.Data(), expected)
}

// TestComplement tests the Complementer interface.
func (t *Tester) TestComplement(test *testing.T) {
	data := []float64{1, -1, 2, -2, 3, -3, 1.5, 7, 0, 1}
	v := t.vec(data)
	anyvec.Complement(v)
	expected := []float64{0, 2, -1, 3, -2, 4, -0.5, -6, 1, 0}
	t.assertClose(test, v.Data(), expected)
}