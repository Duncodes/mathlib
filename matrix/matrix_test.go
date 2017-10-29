package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixNew(t *testing.T) {
	assert := assert.New(t)
	data := New(2, 2, 2, 4, 5, 6)
	_, cols := data.Size()

	assert.Equal(cols, 2, "Cols Should be equal to 2")

}

func TestMatrixZero(t *testing.T) {
	m := Zero(10, 10)
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			value := m.Get(i, j)
			if value != 0 {
				t.Error("Found a none zero elemets in matrix")
			}
		}
	}
}

func TestMatrixGet(t *testing.T) {
	m := New(2, 2, 2, 2, 2, 2)

	for i := 0; i < 2; i++ {
		for j := i; j < 2; j++ {
			value := m.Get(i, j)
			if value != 2 {
				t.Error("Get failed")
			}
		}
	}

}

func TestSize(t *testing.T) {
	m := New(2, 2, 2, 2, 2, 2)

	r, c := m.Size()

	assert := assert.New(t)

	assert.Equal(r, 2, "Should be equal")

	assert.Equal(c, 2, "Should be equal")
}

func TestIndentity(t *testing.T) {
	assert := assert.New(t)

	m := Indentiy(2)

	m1 := New(2, 2, 1, 0, 0, 1)
	assert.Equal(m, m1, "Objects should be equal")

}

func TestGet(t *testing.T) {
	assert := assert.New(t)

	m := New(2, 3, 1, 2, 3, 4, 5, 6)

	k := 1.0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			assert.Equal(k, m.Get(i, j), "Should be equal")
			k++
		}
	}
}

func TestSet(t *testing.T) {
	assert := assert.New(t)

	m := Zero(2, 2)

	m.Set(1, 1, 4)

	assert.Equal(4.0, m.Get(1, 1), "Should be equal")

}

func TestOnes(t *testing.T) {
	assert := assert.New(t)

	m := Ones(3, 3)

	m1 := New(3, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	assert.Equal(m, m1, "Should be equal")

}
