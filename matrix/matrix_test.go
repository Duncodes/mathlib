package matrix

import (
	"testing"
)

func TestMatrixNew(t *testing.T) {
	_, err := New(2, 3, 1, 2, 4, 5, 6)
	if err == nil {
		t.Errorf("Should result to  an error")
	}

	data, _ := New(2, 2, 2, 4, 5, 6)

	if len(data) != 2 {
		t.Errorf("Incorrect size of the new matrix")
	}

}

func TestMatrixZero(t *testing.T) {
	m := Zero(2, 2)
	for _, i := range m {
		for _, j := range i {
			if j != 0 {
				t.Error("Found a non zero vale in matrix")
			}
		}
	}
}

func TestMatrixIndentity(t *testing.T) {
	m := Identity(2)

	if m[0][0] != 1 {
		t.Error("Identity() did not return an identity matrix")
	}

	if m[1][1] != 1 {
		t.Error("Identity() did not return an identity matrix")
	}

	if m[0][1] != 0 {

		t.Error("Identity() did not return an identity matrix")
	}

	if m[1][0] != 0 {
		t.Error("Identity() did not return an identity matrix")
	}

}

func TestSet(t *testing.T) {
	m := Zero(2, 2)
	m.Set(0, 0, 4)

	if m[0][0] != 4 {
		t.Errorf("Set not not working properly")
	}

	if m[0][1] != 0 {
		t.Errorf("Set not not working properly")
	}

	if m[1][0] != 0 {
		t.Errorf("Set not not working properly")
	}

	if m[1][1] != 0 {
		t.Errorf("Set not not working properly")
	}
}
