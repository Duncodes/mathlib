package mathlib

import (
	"testing"
)

func TestReadCSV(t *testing.T) {
	data, err := ReadCSV("testdata/test.csv")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if len(data) != 3 {
		t.Fail()
	}

	_, err = ReadCSVAsFloat("testdata/test.csv")
	if err != nil {
		t.Error("Error with ReadCSVAsFloat64", err)
	}

}
