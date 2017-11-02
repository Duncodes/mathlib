package matrix

import (
	"bytes"
	"errors"
	"strconv"
)

var (
	//ErrOutOfRange your basic index out of range error
	ErrOutOfRange = errors.New("Out of Range")
)

// Matrix represents matrix object
type Matrix struct {
	cols int
	data []float64
}

// String for string matrix formating
func (m *Matrix) String() string {
	buffer := bytes.NewBuffer(nil)

	rows, cols := m.Size()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			buffer.WriteString(strconv.FormatFloat(m.Get(row, col), 'f', -1, 64))
			buffer.WriteRune(' ')
		}
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

// Get gets the data item in the row and col provided,
// Note this is zero besed indexed.Hence the first element of the
// matrix is accessed using Get(0,0)
func (m *Matrix) Get(row, col int) float64 {
	i := (m.cols * row) + col
	return m.data[i]
}

// Size returns the rows and cols of a matrix
func (m *Matrix) Size() (rows, cols int) {
	rows = len(m.data) / m.cols

	cols = m.cols

	return
}

// Set changes the data value at row , col provided.
// Note set has a zero base index ie the first element of the first col
// is accessed via Set(0,0) not Set(1,1)
func (m *Matrix) Set(row, col int, value float64) {
	i := (m.cols * row) + col

	m.data[i] = value

}

// Row returns a given row as a vector
// see https://github.com/Duncodes/mathlib/vector
//func (m *Matrix) Row(row int) vector.Vector {

//}
