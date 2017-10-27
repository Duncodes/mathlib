package stats

import "testing"

func TestPopulationVarience(t *testing.T) {
	data := Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	v, _ := PopulationVariance(data)

	if v != 8.25 {
		t.Errorf("PopulationVariance(data) returned %f but expected %f ", v, 8.25)
	}
}

func TestSampleVarience(t *testing.T) {
	data := Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	v, _ := SampleVariance(data)

	if v != 9.166666666666666 {
		t.Errorf("PopulationVariance(data) returned %f but expected %f ", v, 9.166666666666666)
	}
}

func TestVarienceShort(t *testing.T) {
	data := Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	v, _ := Svariance(data)

	if v != 9.166666666666666 {
		t.Errorf("PopulationVariance(data) returned %f but expected %f ", v, 9.166666666666666)
	}

	data = Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	v, _ = Pvariance(data)

	if v != 8.25 {
		t.Errorf("PopulationVariance(data) returned %f but expected %f ", v, 8.25)
	}

}
