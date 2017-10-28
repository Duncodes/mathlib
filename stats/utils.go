package stats

import (
	"math"
	"math/rand"
	"sort"
	"time"
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

func Sum(d Data) (sum float64, err error) {
	if d.Len() == 0 {
		return math.NaN(), EmptyInput

	}
	for _, n := range d {
		sum += n

	}

	return sum, nil

}

// Samples Data randomly and returns a new Data
func Sample(d Data, lenght int) (Data, error) {
	l := d.Len()
	if l < lenght {
		return nil, SampleSize
	}

	sample := Data{}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < lenght; i++ {
		idx := rand.Intn(l)
		sample = append(sample, d[idx])
	}

	return sample, nil
}
