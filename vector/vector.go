package vector

import (
	"errors"
	"math"
)

var (
	ErrAdd = errors.New("Can't add the vector check and see that they have the same length")
)

// Vector represents a vector
type Vector []float64

// Len returns the lenght of a vector
func (v Vector) Len() int { return len(v) }

// Magnitude computes the Magnitude of the vector
func (v Vector) Magnitude() float64 {
	var square_sum float64
	for _, e := range v {
		square_sum = square_sum + math.Pow(e, 2)
	}

	return math.Pow(square_sum, 0.5)
}

// Is_Unit return if a vector is a unit vector or not
func (v Vector) Is_Unit() bool {
	if v.Magnitude() == 1.0 {
		return true
	}
	return false
}

// Add added two vector
func (v Vector) Add(addvector Vector) (sum Vector, err error) {
	if v.Len() != addvector.Len() {
		return nil, ErrAdd
	}

	for i := 0; i < v.Len(); i++ {
		sum = append(sum, v[i]+addvector[i])
	}

	return
}

// Sub subtracts  two vector
func (v Vector) Sub(addvector Vector) (sum Vector, err error) {
	if v.Len() != addvector.Len() {
		return nil, ErrAdd
	}

	for i := 0; i < v.Len(); i++ {
		sum = append(sum, v[i]-addvector[i])
	}

	return
}
