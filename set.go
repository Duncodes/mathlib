package mathlib

import (
	"fmt"
	"strings"
)

// NewSet create a new set from values
func NewSet(values ...interface{}) Set {
	newset := Set{}
	for _, v := range values {
		newset.Add(v)
	}

	return newset
}

// Set represents a set data type
type Set map[interface{}]bool

// Adds an elemets to the set
func (s Set) Add(i interface{}) {
	s[i] = true
}

func (s Set) Remove(i interface{}) {
	delete(s, i)
}

// Contains returns  true if an element is contained in set
func (s Set) Contains(i interface{}) bool {
	return s[i]
}

// Len returns length of the set
func (s Set) Len() int {
	return len(s)
}

// String returns a set as a string
func (s Set) String() string {
	var values []string

	for k := range s {
		values = append(values, fmt.Sprintf("%v", k))
	}

	return "{ " + strings.Join(values, " , ") + " }"
}

// IsSubsetOf returns true if set contains another set
func (s Set) IsSubsetOf(set Set) bool {

	// this set is a subset of set provided
	for k := range s {
		if !set.Contains(k) {
			return false
		}
	}

	return true
}

// Union joins two sets together
func (s Set) Union(set Set) Set {
	u := Set{}

	for s := range s {
		u.Add(s)
	}

	for s := range set {
		u.Add(s)
	}

	return u
}

// Intersection return a set of  those elemets that are in both sets
func (s Set) Intersection(set Set) Set {
	intersection := Set{}

	for k := range s {
		if set.Contains(k) {
			intersection.Add(k)
		}
	}

	return intersection
}

// Difference returns a set of elements that are in the object set and not in provided set
func (s Set) Difference(set Set) Set {
	diff := Set{}

	for k := range s {
		if !set.Contains(k) {
			diff.Add(k)
		}
	}

	return diff
}

func (s Set) Cardinality() int {
	return s.Len()
}
