package main

import (
	"fmt"

	"github.com/Duncodes/mathlib/stats"
)

// This is something i always wanted to test why n-1 in SampleStandardDeviation.
// Well this proves it. At avarange this is always a better eximation of the population StandardDeviation
func main() {
	s := stats.Data{1, 2, 7, 2, 1, 2, 4, 6, 2, 8, 8, 8, 8, 8, 5, 7, 3, 2, 1, 5, 8, 9, 5, 6, 3, 4, 5, 6, 7, 7, 9, 4, 5, 6, 2, 3, 4, 1, 6, 4, 5, 6, 2, 3, 4, 1, 4, 4, 2, 3, 4, 2, 4, 6, 7, 4, 2, 5, 6, 2, 4, 7, 3, 4, 2, 4, 5, 3, 4, 2}

	std, _ := s.PopulationStandardDeviation()

	var inter int = 1000
	var as_sample_sum float64 = 0
	var as_pop_sum float64 = 0

	for i := 0; i < inter; i++ {
		sample, _ := stats.Sample(s, 5)
		//fmt.Println(sample)
		d, _ := sample.SampleStandardDeviation()
		as_sample_sum += d

		sample1, _ := stats.Sample(s, 5)
		d1, _ := sample1.PopulationStandardDeviation()
		as_pop_sum += d1
	}

	fmt.Println("1 => The sample std averange ", as_sample_sum/float64(inter))
	fmt.Println("2 => The eximate population  std averange ", as_pop_sum/float64(inter))
	fmt.Println("3 => The acutal std ", std)

	fmt.Println("Note that on average 1 is always a better approximation of the true std than 2")

}
