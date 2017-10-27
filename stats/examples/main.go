package main

import (
	"fmt"

	"github.com/Duncodes/mathlib/stats"
)

func main() {
	data := stats.Data{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	mean, _ := data.Mean()
	median, _ := data.Median()
	mode, _ := data.Mode()
	variance, _ := data.Variance()
	std, _ := data.StandardDeviation()

	s, _ := stats.Sstd(data)

	fmt.Println("Mean: ", mean)

	fmt.Println("Median ", median)

	fmt.Println("Mode ", mode)

	fmt.Println("Variance", variance)

	fmt.Println("StandardDeviation", std)

	fmt.Println("StandardDeviation Sample", s)

	d, _ := data.SampleZscore(6)
	d1, _ := data.Zscore(6)

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
}
