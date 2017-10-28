package stats

import (
	"testing"
)

func TestStandardDeviation(t *testing.T) {
	data := Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	v, _ := PopulationStandardDeviation(data)

	if v != 2.8722813232690143 {
		t.Errorf("PopulationStandardDeviation(data) returned %f but expected %f ", v, 2.8722813232690143)
	}

	v, _ = SampleStandardDeviation(data)

	if v != 3.0276503540974917 {
		t.Errorf("SampleStandardDeviation(data) returned %f but expected %f ", v, 3.0276503540974917)
	}

	// short

	v, _ = Ssdev(data)

	if v != 3.0276503540974917 {
		t.Errorf("Ssdev(data) returned %f but expected %f ", v, 3.0276503540974917)
	}

	v, _ = Psdev(data)

	if v != 2.8722813232690143 {
		t.Errorf("Psdev returned %f but expected %f ", v, 2.8722813232690143)
	}

}

func TestZscore(t *testing.T) {
	data := Data{3, 3, 5, 5, 4}

	z, _ := SampleZscore(data, 6)
	if z != 2 {
		t.Errorf("SampleZscore(data) failed to produce collect result")
	}

	z, _ = Zscore(data, 4)
	if z != 0 {
		t.Errorf("Zscore(data) failed to produce collect result")
	}

}

func TestSampleStandardError(t *testing.T) {
	data := Data{3, 3, 5, 5, 4}
	r, _ := SampleStandardError(data)

	if r != 0.4472135954999579 {
		t.Errorf("SampleStandardError(data) expected %f got %f", 0.4472135954999579, r)
	}
}
