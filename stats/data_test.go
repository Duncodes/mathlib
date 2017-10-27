package stats

import (
	"reflect"
	"testing"
)

var data = Data{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestDataHelperMethods(t *testing.T) {
	f := data.Get(1)
	if f != 2 {
		t.Errorf("Get(1) returned  %f but expected %f", f, 2)
	}

	l := data.Len()
	if l != 9 {
		t.Errorf("Len() returned %d but expected %d", l, 9)
	}

	min, _ := data.Min()
	if min != 1 {
		t.Errorf("Min() returned %f but expected %f", min, 1.0)
	}

	max, _ := data.Max()
	if max != 9 {
		t.Errorf("Max() returned %f but expected %f", max, 9.0)
	}

}

func TestDataMethods(t *testing.T) {
	// create test data
	var tdata = Data{1, 2, 3, 4, 5}

	// test mean
	mean, _ := tdata.Mean()
	if mean != 3 {
		t.Errorf("Mean() returned %f but expected %f", mean, 3.0)
	}

	// test medium
	medium, _ := tdata.Median()
	if medium != 3 {
		t.Errorf("Median() returned %f but expected %f", medium, 3.0)
	}

	modef := Data{1, 2, 3, 8, 2, 3.0}
	// test the mode
	mode, _ := modef.Mode()
	if !reflect.DeepEqual(mode, []float64{2, 3}) {
		t.Errorf("Mode() returned %f but expected  %f", mode, []float64{2, 3})
	}

	sum, _ := tdata.Sum()
	if sum != 15 {
		t.Errorf("Sum() returned %f but expected %f ", sum, 15)
	}

}
