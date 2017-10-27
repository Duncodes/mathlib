package stats

import "math"

// Median computes the slice median
func Median(data Data) (median float64, err error) {
	// create a sorted copy
	c := sortCopy(data)

	l := len(c)
	if l == 0 {
		return math.NaN(), EmptyInput
	} else if l%2 == 0 {
		median, err = Mean(c[l/2-1 : l/2+1])
	} else {
		median = float64(c[l/2])
	}

	return
}
