package stats

import "math"

// Range computes range
func Range(d Data) (float64, error) {
	if d.Len() == 0 {
		return math.NaN(), EmptyInput
	}

	c := sortCopy(d)

	return c[d.Len()-1] - c[0], nil
}
