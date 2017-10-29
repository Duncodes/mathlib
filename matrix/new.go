package matrix

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

// Indentiy create an Indentiy matrix
func Indentiy(order int) *Matrix {
	m := New(order, order)
	for i := 0; i < order; i++ {
		m.Set(i, i, 1)
	}

	return m
}
