// Package matrix provides mathematical impremetation  of matrix
package matrix

import (
	"errors"
	"fmt"
)

type Matrix [][]float64

// Set sets a value to row and col
// returns error if row, col are out of range
// It will override the inital value
func (m Matrix) Set(row int, col int, value float64) error {
	m[row][col] = value
	fmt.Println(m)
	return nil
}

// Size returns the  dimensions of the matrix
func (m *Matrix) Size() (row, col int) {
	col = len(*m)
	//row = len([][]m)

	return
}

//New creates a new matrix with inital values provided
//The values provide should be equal to the rows*cols. Otherwise
//A matrix won't be formed.Values should alway be inform of rows.
//If all the values are not provided it will always return a Zero matrix
func New(rows int, cols int, values ...float64) (Matrix, error) {
	var m Matrix

	valueslen := rows * cols // total elements of a matrix

	if len(values) < valueslen {
		return Zero(rows, cols), errors.New("Must provide all the values ")
	}

	for i := 0; i < valueslen; i = i + cols {
		m = append(m, values[i:cols+i]) // slice only what you need
	}
	return m, nil
}

// Zero creates a zero matrix. Given it rows and col
func Zero(rows, cols int) Matrix {
	m := new(Matrix)

	c := make([]float64, cols)
	for i := 0; i < rows; i++ {
		*m = append(*m, c)
	}

	return *m
}

// Identity return the identity matrix of a given order
func Identity(order int) Matrix {
	var m Matrix
	m = Zero(order, order)
	for i := 0; i < order; i++ {
		for j := 0; j < order; j++ {
			if i == j {
				m.Set(i, j, 1)
			}
		}
	}

	return m
}
