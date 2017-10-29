package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixFromArray(t *testing.T) {
	assert := assert.New(t)

	m := MatrixFromArray([][]float64{{1, 1}, {1, 2}})

	m2 := New(2, 2, 1, 1, 1, 2)

	assert.Equal(m, m2, "m and m2 Should be equal")

}

func TestArray(t *testing.T) {
	assert := assert.New(t)

	a1 := [][]float64{{1, 2, 4, 5}, {6, 7, 8, 9}}
	a2 := [][]float64{{5, 6}, {4, 3}}
	a3 := [][]float64{{9, 8, 1, 3}, {6, 5, 5, 3}, {7, 8, 2, 7}, {1, 3, 6, 7}}

	m1 := MatrixFromArray(a1)
	m2 := New(2, 2, 5, 6, 4, 3)
	m3 := New(4, 4, 9, 8, 1, 3, 6, 5, 5, 3, 7, 8, 2, 7, 1, 3, 6, 7)

	m4 := m3.Array()
	m5 := m2.Array()

	m6 := m1.Array()
	assert.Equal(m5, a2, "Should be equal")
	assert.Equal(a3, m4, "Should be equal")

	assert.Equal(a1, m6, "Should be equal")
}
