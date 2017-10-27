package stats

import (
	"errors"
	"math"
)

// std computes both sample and  population
func std(d Data, sample bool) (std float64, err error) {
	if sample == true {
		variance, _ := d.SampleVariance()
		return math.Pow(variance, 0.5), nil
	}
	variance, _ := d.PopulationVariance()
	return math.Pow(variance, 0.5), nil

}

// StandardDeviation computes the population StandardDeviation
func StandardDeviation(d Data) (float64, error) {
	return PopulationStandardDeviation(d)
}

// StandardDeviation computes the population StandardDeviation
func PopulationStandardDeviation(d Data) (float64, error) {
	return std(d, false)
}

// SampleStandardDeviation computes the sample StandardDeviation
func SampleStandardDeviation(d Data) (float64, error) {
	return std(d, true)
}

// Sstd is a short for SampleStandardDeviation
func Sstd(d Data) (float64, error) {
	return SampleStandardDeviation(d)
}

// Pstd is a short for PopulationStandardDeviation
func Pstd(d Data) (float64, error) {
	return PopulationStandardDeviation(d)
}

// Zscore self
func Zscore(d Data, x float64) (z float64, err error) {
	m, err := Mean(d)

	s, _ := d.StandardDeviation()

	if s == 0 {
		return math.NaN(), errors.New("Not defined")
	}

	z = float64(x-m) / s
	return
}

// SampleZscore self
func SampleZscore(d Data, x float64) (z float64, err error) {
	m, err := Mean(d)
	s, _ := d.SampleStandardDeviation()

	if s == 0 {
		return math.NaN(), errors.New("Not defined")
	}
	z = float64(x-m) / s

	return

}

func SampleStandardError(d Data) (float64, error) {
	s, _ := d.SampleStandardDeviation()
	if d.Len() == 0 {
		return math.NaN(), EmptyInput
	}

	return s / math.Pow(float64(d.Len()), 0.5), nil
}
