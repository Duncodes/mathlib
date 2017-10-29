package stats

// Data represents a collection of float64 values. In which
// operations are done
type Data []float64

// Get items in slice
func (d Data) Get(i int) float64 { return d[i] }

// Len return length of slice
func (d *Data) Len() int { return len(*d) }

// Mean Computes the of the data set
func (d *Data) Mean() (float64, error) { return Mean(*d) }

// Sum returns the sum of all the values of data
func (d *Data) Sum() (float64, error) { return Sum(*d) }

// Median returns the mid item in the data when sorted
func (d *Data) Median() (float64, error) { return Median(*d) }

// Mode returns the mode of the data
func (d *Data) Mode() ([]float64, error) { return Mode(*d) }

// Max returns the maximux values in data
func (d *Data) Max() (float64, error) { return Max(*d) }

// Min return the minimum item in data
func (d *Data) Min() (float64, error) { return Min(*d) }

// Variance computes the Population Variance of  data
func (d *Data) Variance() (float64, error) { return Variance(*d) }

// PopulationVariance return the Population Variance of data
func (d *Data) PopulationVariance() (float64, error) {
	return d.Variance()
}

// SampleVariance returns sample Variance of data
func (d *Data) SampleVariance() (float64, error) {
	return SampleVariance(*d)
}

// StandardDeviation computes the Population StandardDeviation
func (d *Data) StandardDeviation() (float64, error) {
	return StandardDeviation(*d)
}

// PopulationStandardDeviation computes the Population StandardDeviation
func (d *Data) PopulationStandardDeviation() (float64, error) {
	return PopulationStandardDeviation(*d)
}

// SampleStandardDeviation computes the sample StandardDeviation
func (d *Data) SampleStandardDeviation() (float64, error) {
	return SampleStandardDeviation(*d)
}

// Zscore computes the number of StandardDeviation from the mean using Population
func (d *Data) Zscore(x float64) (float64, error) {
	return Zscore(*d, x)
}

// SampleZscore computes the number of StandardDeviation from the mean using sample
func (d *Data) SampleZscore(x float64) (float64, error) {
	return SampleZscore(*d, x)
}

// Range return the data range
func (d *Data) Range() (float64, error) { return Range(*d) }

// SampleStandardError computes sample StandardError
func (d *Data) SampleStandardError() (float64, error) { return SampleStandardError(*d) }
