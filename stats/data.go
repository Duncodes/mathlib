//
package stats

import "math"

type Data []float64

// Get items in slice
func (d Data) Get(i int) float64 { return d[i] }

// Len return length of slice
func (d Data) Len() int { return len(d) }

// Min return min
//func (f Data) Min() (float64, error){ return }

func (d Data) Mean() (float64, error) { return Mean(d) }

func (d Data) Sum() (float64, error) { return sum(d) }

func (d Data) Median() (float64, error) { return Median(d) }

func (d Data) Mode() ([]float64, error) { return Mode(d) }

func (d Data) Max() (float64, error) { return Max(d) }

func (d Data) Min() (float64, error) { return Min(d) }

func (d Data) Variance() (float64, error) { return Variance(d) }

func (d Data) PopulationVariance() (float64, error) {
	return d.Variance()
}

func (d Data) SampleVariance() (float64, error) {
	return SampleVariance(d)
}

func (d Data) StandardDeviation() (float64, error) {
	return StandardDeviation(d)
}

func (d Data) PopulationStandardDeviation() (float64, error) {
	return PopulationStandardDeviation(d)
}

func (d Data) SampleStandardDeviation() (float64, error) {
	return SampleStandardDeviation(d)
}

func (d Data) Zscore(x float64) (float64, error) {
	return Zscore(d, x)
}

func (d Data) SampleZscore(x float64) (float64, error) {
	return SampleZscore(d, x)
}

func (d Data) Range() (float64, error) { return Range(d) }

func (d Data) SampleStandardError() (float64, error) { return SampleStandardError(d) }

func sum(i Data) (sum float64, err error) {
	if i.Len() == 0 {
		return math.NaN(), EmptyInput

	}

	// Add em up
	for _, n := range i {
		sum += n

	}

	return sum, nil

}
