package matrix_test

import (
	"fmt"

	"github.com/Duncodes/mathlib/matrix"
)

func ExampleNew() {
	m := matrix.New(1, 2, 4, 3)

	fmt.Println(m)
	// Output: 4 3

}

func ExampleZero() {
	m := matrix.Zero(1, 3)
	fmt.Println(m)

	// Output: 0 0 0
}

// This creates a new Matrix from	[][]float64
func ExampleNewFromArray() {
	m := matrix.NewFromArray([][]float64{{1, 2}, {1, 2}})

	fmt.Println(m)

}
