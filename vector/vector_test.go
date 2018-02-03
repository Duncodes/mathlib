package vector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector(t *testing.T) {
	v := Vector{1, 2, 3, 4}

	// test len
	if v.Len() != 4 {
		t.Errorf("Vector.Len() should be %f", v.Len())
	}

	assert := assert.New(t)

	v = Vector{3, -5, 10}

	m := v.Magnitude()

	assert.Equal(math.Pow(134, 0.5), m, "Should be equal")

	unit := Vector{1, 0, 0}

	assert.True(unit.Is_Unit(), "Should be true")

	unit = Vector{1, 0, 2}

	unit2 := Vector{1, 0, 0}

	assert.False(unit.Is_Unit(), "Should be false")

	sum := Vector{2, 0, 2}

	sum1, err := unit.Add(unit2)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(sum, sum1, "Should be equal")

	vec1 := Vector{1, 2}

	new_vec := New(1, 2)

	assert.Equal(vec1, new_vec, "Should be equal")

	vec2 := vec1.Scalar(2)

	vec3 := Vector{2, 4}

	assert.Equal(vec2, vec3, "Should be Equal")
}
