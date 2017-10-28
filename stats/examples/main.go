package main

import (
	"fmt"

	"github.com/Duncodes/mathlib/stats"
)

func main() {
	data := stats.Data{3, 3, 5, 5, 4}

	mean, _ := data.Mean()
	median, _ := data.Median()
	mode, _ := data.Mode()
	variance, _ := data.SampleVariance()
	std, _ := data.StandardDeviation()

	s, _ := stats.Ssdev(data)

	fmt.Println("Mean: ", mean)

	fmt.Println("Median ", median)

	fmt.Println("Mode ", mode)

	fmt.Println("Variance", variance)

	fmt.Println("StandardDeviation", std)

	fmt.Println("StandardDeviation Sample", s)

	d, _ := data.SampleZscore(6)
	d1, _ := data.Zscore(4)

	fmt.Println("Sample ZScore ", d)
	fmt.Println("ZScore ", d1)
	max, _ := data.Max()
	min, _ := data.Min()
	sum, _ := data.Sum()
	r, _ := data.Range()
	sserror, _ := data.SampleStandardError()

	fmt.Println("Max ", max)
	fmt.Println("Min ", min)
	fmt.Println("Sum ", sum)
	fmt.Println("Range", r)
	fmt.Println("Std Error", sserror)

	s1 := stats.Data{1, 2, 3, 4}
	s2 := stats.Data{5, 6, 7, 8}

	r, _ = stats.Correlation(s1, s2)
	fmt.Println("Correlation ", r)
}
