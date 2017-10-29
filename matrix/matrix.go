package matrix

import (
	"bytes"
	"strconv"
)

// Matrix represents
type Matrix struct {
	cols int
	data []float64
}

// String imprements a stringer interface
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

// New returns a new matrix provided given the order . Ie rows and cols.
// Values are the values of a martix as cols, If less values are provided than the
// order zeros are used to fill the matrix .
func New(rows, cols int, values ...float64) *Matrix {
	if len(values) == 0 {
		return &Matrix{
			cols: cols,
			data: make([]float64, rows*cols),
		}
	}
	data := make([]float64, rows*cols)

	copy(data, values)
	return &Matrix{
		cols: cols,
		data: data,
	}

}

// Zero creates a zero matrix
func Zero(rows, cols int) *Matrix {
	return New(rows, cols)
}

// Ones create a ones matrix
func Ones(rows, cols int) *Matrix {
	m := New(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m.Set(i, j, 1)
		}
	}

	return m
}

// Size returns the rows and cols of a matrix
func (m *Matrix) Size() (rows, cols int) {
	rows = len(m.data) / m.cols

	cols = m.cols

	return
}

// Indentiy create an Indentiy matrix
func Indentiy(order int) *Matrix {
	m := New(order, order)
	for i := 0; i < order; i++ {
		m.Set(i, i, 1)
	}

	return m
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
