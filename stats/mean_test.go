package stats

import (
	"testing"
)

func TestMean(t *testing.T) {
	var d1 = Data{0, 0, 0, 0, 0, 0, 0}
	var d2 = Data{-1, -2, -3, -4, -5}
	var d3 = Data{-1, -2, 0, 1, 2}

	m1, _ := d1.Mean()
	m2, _ := d2.Mean()
	m3, _ := d3.Mean()

	if m3 != 0 {
		t.Errorf("Mean() failed %f, %f ,  %f", m1, m2, m3)
	}
}
