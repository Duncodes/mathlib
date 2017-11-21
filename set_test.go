package mathlib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	assert := assert.New(t)

	set := NewSet(1, 2, 3, "numbers")

	assert.True(set.Contains(1), "Should contain 1")

	set.Add("all good")

	assert.True(set.Contains("all good"), "Should contain all good")

	_ = fmt.Sprintln(set)

	assert.Len(set, set.Len(), "Should have the same length")

	set.Remove(1)

	assert.False(set.Contains(1), "Should be false, Should have been removed")

	assert.Equal(set.Cardinality(), set.Len(), "Cardinality should be len")
}

func TestSetOperations(t *testing.T) {
	assert := assert.New(t)

	set := NewSet(1, 2, 3)

	set1 := NewSet(1, 2)

	assert.True(set1.IsSubsetOf(set))

	// union

	set = NewSet("cows", "dogs")

	set1 = NewSet("cows", "cats", "rats")

	uset := set.Union(set1)

	u := NewSet("cows", "cats", "rats", "dogs")
	assert.Equal(uset, u, "Should be a union of two sets")

	iset := set.Intersection(set1)

	i := NewSet("cows")
	assert.Equal(iset, i, "Should be equal")

	diff := NewSet("rats", "cats")

	d := set1.Difference(set)

	assert.Equal(diff, d, "Should be the diff")
}
