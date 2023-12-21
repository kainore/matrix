package matrix

import (
	"fmt"
	"sync"
)

// Matrix represents a 2D matrix.
type Matrix struct {
	rows, cols int
	data       [][]float64
}

// New creates a new matrix with the specified dimensions.
func New(rows, cols int) *Matrix {
	m := &Matrix{rows: rows, cols: cols, data: make([][]float64, rows)}
	for i := range m.data {
		m.data[i] = make([]float64, cols)
	}
	return m
}

// Set sets the value at the specified row and column.
func (m *Matrix) Set(row, col int, value float64) {
	m.data[row][col] = value
}

// Get retrieves the value at the specified row and column.
func (m *Matrix) Get(row, col int) float64 {
	return m.data[row][col]
}

// Multiply multiplies two matrices concurrently using Goroutines.
func Multiply(a, b *Matrix) *Matrix {
	if a.cols != b.rows {
		panic("Matrix dimensions do not match for multiplication")
	}

	result := New(a.rows, b.cols)
	var wg sync.WaitGroup

	for i := 0; i < a.rows; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < b.cols; j++ {
				var sum float64
				for k := 0; k < a.cols; k++ {
					sum += a.Get(i, k) * b.Get(k, j)
				}
				result.Set(i, j, sum)
			}
		}(i)
	}

	wg.Wait()
	return result
}

// Print displays the matrix.
func (m *Matrix) Print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%.2f\t", m.data[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
