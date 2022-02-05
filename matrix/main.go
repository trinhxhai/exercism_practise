package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	data [][]int
	n, m int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	n, m := 0, 0
	n = len(rows)
	if n > 0 {
		m = len(strings.Split(rows[0], " "))
	}
	data := make([][]int, n)

	for i, row := range rows {

		cols := strings.Split(strings.TrimSpace(row), " ")

		s_row := make([]int, m)

		for j, col := range cols {
			res, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			s_row[j] = res
		}

		data[i] = s_row
	}

	return &Matrix{
		data: data,
		n:    n,
		m:    m,
	}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	res := make([][]int, m.m)
	for i := 0; i < m.m; i++ {
		res[i] = make([]int, m.n)
		for j := 0; j < m.n; j++ {
			res[i][j] = m.data[j][i]
		}
	}
	return res
}

func (m *Matrix) Rows() [][]int {
	return m.data
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= m.n || col < 0 || col >= m.m {
		return false
	}
	m.data[row][col] = val
	return true
}

func main() {
	input := "1 2\n10 20"

	mx, _ := New(input)
	fmt.Printf("%v\n", mx)
	fmt.Printf("%v\n", mx.Cols())
	fmt.Printf("%v\n", mx.Rows())

}
