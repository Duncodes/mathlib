package stats

import (
	"math"
	"sort"
)

func sortCopy(data Data) Data {
	c := make(Data, data.Len())
	copy(c, data)
	sort.Float64s(c)
	return c
}

func Max(d Data) (float64, error) {
	if d.Len() == 0 {
		return math.NaN(), EmptyInput
	}

	max := d.Get(0)

	for i := 1; i < d.Len(); i++ {
		if d.Get(i) > max {
			max = d.Get(i)
		}
	}

	return max, nil
}

func Min(d Data) (float64, error) {
	if d.Len() == 0 {
		return math.NaN(), EmptyInput
	}

	min := d.Get(0)

	for i := 1; i < d.Len(); i++ {
		if d.Get(i) < min {
			min = d.Get(i)
		}
	}

	return min, nil

}

func Sum(d Data) (sum float64) {
	if d.Len() == 0 {
		return math.NaN()

	}

	// Add em up
	for _, n := range d {
		sum += n

	}

	return sum

}
