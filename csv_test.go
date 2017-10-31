package mathlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSV(t *testing.T) {

	assert := assert.New(t)
	data, err := ReadCSV("testdata/test.csv")
	assert.NoError(err, "Should be nill")

	if len(data) != 3 {
		t.Fail()
	}
	data, err = ReadCSV("testdata/test_nil.csv")
	assert.Error(err, "Should return an error")

	data, err = ReadCSV("testdata/test_error.csv")
	assert.Error(err, "Should return err")

	_, err = ReadCSVAsFloat("testdata/test.csv")
	assert.NoError(err, "Should not return an error")

	_, err = ReadCSVAsFloat("testdata/test_nil.csv")
	assert.Error(err, "Should return an error")

	_, err = ReadCSVAsFloat("testdata/test_error.csv")
	assert.Error(err, "Should return an error")

}
