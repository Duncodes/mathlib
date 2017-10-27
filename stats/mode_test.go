package stats

import (
	"reflect"
	"testing"
)

func TestMode(t *testing.T) {
	var data = Data{1, 2, 3, 3, 24, 4}

	mode, _ := Mode(data)

	if !reflect.DeepEqual(mode, []float64{3}) {
		t.Errorf("Mode(data) returned %f but expected %f ", mode, []float64{3})
	}
}
