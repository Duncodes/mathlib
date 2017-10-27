package stats

func Mode(i Data) (mode []float64, err error) {
	l := i.Len()
	if l == 0 {
		return nil, EmptyInput
	}

	c := sortCopy(i)

	mode = make([]float64, 5)
	cnt, maxCnt := 1, 1
	for j := 1; j < l; j++ {
		switch {
		case c[j] == c[j-1]:
			cnt++
		case cnt == maxCnt && maxCnt != 1:
			mode = append(mode, c[j-1])
			cnt = 1
		case cnt > maxCnt:
			mode = append(mode[:0], c[j-1])
			maxCnt, cnt = cnt, 1
		default:
			cnt = 1

		}

	}
	switch {
	case cnt == maxCnt:
		mode = append(mode, c[l-1])
	case cnt > maxCnt:
		mode = append(mode[:0], c[l-1])
		maxCnt = cnt

	}

	// Since length must be greater than 1,
	// check for slices of distinct values
	if maxCnt == 1 {
		return Data{}, nil

	}

	return mode, nil
}
