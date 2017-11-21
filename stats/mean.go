package stats

import "math"

// Mean get the average of a slice of number
func Mean(i Data) (float64, error) {
	if i.Len() == 0 {
		return math.NaN(), EmptyInput
	}

	sum, _ := i.Sum()
	return sum / float64(i.Len()), nil
}
