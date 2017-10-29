package matrix

func MatrixFromArray(data [][]float64) *Matrix {
	rows := len(data)

	cols := len(data[0])

	m := New(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m.Set(i, j, data[i][j])
		}
	}

	return m
}

func (m *Matrix) Array() [][]float64 {
	rows, cols := m.Size()

	a := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		a[i] = make([]float64, cols)

		for j := 0; j < cols; j++ {
			a[i][j] = m.Get(i, j)
		}
	}

	return a
}
