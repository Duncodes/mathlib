package mathlib

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ReadCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	// f is a Reader
	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// ReadCSVAsFloat returns the values of csv as floats
// TODO runtime for very large csv file ???
func ReadCSVAsFloat(filename string) ([][]float64, error) {
	var data [][]float64

	records, err := ReadCSV(filename)
	if err != nil {
		return nil, err
	}

	for _, r := range records {
		var f []float64
		for _, v := range r {
			fl, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, err
			}
			f = append(f, fl)
		}
		data = append(data, f)
	}

	return data, nil
}
