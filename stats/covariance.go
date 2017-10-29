package stats

import "math"

// Correlation computes the Pearson
func Correlation(d, d1 Data) (float64, error) {
	l1, l2 := d.Len(), d1.Len()

	if l1 != l2 {
		return math.NaN(), SizeErr
	}

	if l1 == 0 || l2 == 0 {
		return math.NaN(), EmptyInput
	}
	var sum float64

	for i := 0; i < d.Len(); i++ {
		z1, _ := d.Zscore(d[i])
		z2, _ := Zscore(d1, d1[i])
		sum += z1 * z2
	}

	return sum / float64(d.Len()), nil
}

// Pearson computes the pearson moment coef
func Pearson(d1, d2 Data) (float64, error) {
	return Correlation(d1, d2)
}
