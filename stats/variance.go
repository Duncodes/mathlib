package stats

import (
	"math"
)

// variance computes the variance of smaple if sample is true else population
func variance(data Data, sample bool) (variance float64, err error) {
	if data.Len() == 0 {
		return math.NaN(), EmptyInput
	}
	is_sample := 0
	if sample {
		is_sample = 1
	}
	mean, err := data.Mean()
	for _, n := range data {
		variance += math.Pow((float64(mean) - float64(n)), 2)
	}

	variance = variance / float64((data.Len())-is_sample)
	return
}

// Variance ccomputes the population variance
func Variance(data Data) (float64, error) {
	return PopulationVariance(data)
}

// PopulationVariance computes the population variance
func PopulationVariance(data Data) (float64, error) {
	return variance(data, false)
}

// Pvariance computes the population variance
func Pvariance(data Data) (float64, error) {
	return variance(data, false)
}

// SampleVariance computes the sample variance
func SampleVariance(data Data) (float64, error) {
	return variance(data, true)
}

// Svariance computes the sample variance
func Svariance(data Data) (float64, error) {
	return variance(data, true)
}
