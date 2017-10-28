package stats

import "testing"

func TestCorrelation(t *testing.T) {
	// test correlation
	d1 := Data{1, 2, 3, 4}
	d2 := Data{1, 2, 3}
	d3 := Data{5, 6, 7, 8}
	d4 := Data{}

	// test zero size error
	_, err := Correlation(d4, d2)
	if err == nil {
		t.Errorf("Correlation() should have returned an error")
	}

	r, _ := Correlation(d1, d3)
	if r != 1 {
		t.Errorf("Correlation() gave  us %f but we expected %f ", r, 1.0)
	}

	// test pearsons returns Correlation
	d5 := Data{1, 2, 3}
	c, _ := Correlation(d2, d5)
	r, _ = Pearson(d2, d5)
	if c != r {
		t.Errorf("Pearson() did not return Correlation() ")
	}
}
