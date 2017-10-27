package vector

import "testing"

func TestVector(t *testing.T) {
	v := Vector{1, 2, 3, 4}

	// test len
	if v.Len() != 4 {
		t.Errorf("Vector.Len() should be %f", v.Len())
	}
}
