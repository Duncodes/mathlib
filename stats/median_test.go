package stats

import "testing"

func TestMedian(t *testing.T) {
	var tdata = Data{1, 2, 3, 4, 5}

	median, _ := Median(tdata)

	if median != 3 {
		t.Errorf("Median(data) returned %f but expected %f", median, 3)
	}

	tdata = Data{1, 3, 4, 5}

	median, _ = Median(tdata)

	if median != 3.5 {
		t.Errorf("Median(data) returned %f but expected %f", median, 3.5)
	}

}
