package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	m := Matrix{}
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		cols := strings.Split(strings.TrimSpace(l), " ")
		line := []int{}
		for _, c := range cols {
			val, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			line = append(line, val)
		}
		if len(line) == 0 || len(m) > 0 && len(m[i-1]) != len(line) {
			return nil, errors.New("something wrong")
		}
		m = append(m, line)
	}
	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	result := [][]int{}
	for k := 0; k < len(m[0]); k++ {
		line := []int{}
		for i := range m {
			line = append(line, m[i][k])
		}
		result = append(result, line)
	}
	return result
}

func (m Matrix) Rows() [][]int {
	result := [][]int{}
	for _, v := range m {
		line := []int{}
		line = append(line, v...)
		result = append(result, line)
	}
	return result
}

func (m Matrix) Set(row, col, val int) bool {
	if row >= 0 && col >= 0 && len(m) > row && len(m[row]) > col {
		m[row][col] = val
		return true
	}
	return false
}
