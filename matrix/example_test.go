package matrix_test

import (
	"fmt"

	"github.com/Duncodes/mathlib/matrix"
)

func ExampleNew() {
	m, err := matrix.New(2, 2, 4, 3, 3, 4)
	if err != nil {
		// handle error
	}

	fmt.Println(m)
	//Output: [[4 3] [3 4]]
}

func ExampleZero() {
	m := matrix.Zero(2, 2)
	fmt.Println(m)
	// Output: [[0 0] [0 0]]
}
