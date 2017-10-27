package stats

import "testing"

func TestRange(t *testing.T) {
	data := Data{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 0, 6, 5, 4, 43, 53, 4, 3, 43}

	r, _ := Range(data)

	if r != 53 {
		t.Errorf("Range(data) is block")
	}

}
