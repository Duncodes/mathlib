package vector

// Vector represents a vector
type Vector []float64

// Len returns the lenght of a vector
func (v Vector) Len() int { return len(v) }
