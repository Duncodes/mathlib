package stats

import (
	"log"
	"testing"
)

func TestMax(t *testing.T) {
	data := Data{1, 2, 5, 3, 4}

	max, _ := Max(data)
	if max != 5 {
		t.Errorf("Max(data) returned %f but expected %f", max, 5)
	}
}

func TestMin(t *testing.T) {
	data := Data{1, 2, 5, 3, 4}

	min, _ := Min(data)

	if min != 1 {
		t.Errorf("Min(data) returned %f but expected %f ", min, 1)
	}

}

func TestSum(t *testing.T) {

	data := Data{1, 2, 3, 4, 5, 6}

	sum := Sum(data)
	if sum != 21 {
		t.Errorf("Sum(data) returned %f but expected %f ", sum, 1)
	}
}

func TestSample(t *testing.T) {
	data := Data{1, 2, 3, 4, 5, 6, 7, 8, 9}

	newsample, _ := Sample(data, 4)

	if (newsample.Len()) != 4 {
		t.Errorf("Sample(data) failed")
	}
	log.Println(newsample)
}
