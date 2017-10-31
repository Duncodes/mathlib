package mathlib_test

import (
	"fmt"

	"github.com/Duncodes/mathlib"
)

func ExampleNewSet() {
	set := mathlib.NewSet(1, 2, 3, 3, 4)

	fmt.Println(set)

}
