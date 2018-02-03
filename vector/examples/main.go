package main

import (
	"fmt"

	"github.com/Duncodes/mathlib/vector"
)

func main() {
	//vector := []float64{1, 2, 3}

	vector := vector.Vector{1, 2, 3}

	fmt.Println(vector.Len())
}
