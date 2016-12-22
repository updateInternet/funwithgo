package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	matrix, err := newmatrix(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", matrix.AbsoluteDiagonalDifference())
}

type matrix struct {
	Dimension int
	Data      [][]int
}

func newmatrix(r io.Reader) (*matrix, error) {
	m := new(matrix)
	// Fscanf treats new line as spaces,which is what we want for this excersie
	_, err := fmt.Fscanf(r, "%d", &m.Dimension)
	if err != nil {
		return nil, err
	}

	m.Data = make([][]int, m.Dimension)

	for i := range m.Data {
		m.Data[i] = make([]int, m.Dimension)
	}

	for y := 0; y < m.Dimension; y++ {
		for x := 0; x < m.Dimension; x++ {
			_, err := fmt.Fscanf(r, "%d", &m.Data[y][x])
			if err != nil {
				return nil, err
			}
		}
	}

	return m, nil
}

func (m *matrix) Print() {
	for y := 0; y < m.Dimension; y++ {
		for x := 0; x < m.Dimension; x++ {
			fmt.Printf("%d", m.At(x, y))

			if x != m.Dimension-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func (m *matrix) At(x, y int) int {
	return m.Data[y][x]
}

func (m *matrix) Diagonal() int {
	sum := 0
	for i := 0; i < m.Dimension; i++ {
		sum += m.At(i, i)
	}
	return sum
}

func (m *matrix) ReverseDiagonal() int {
	sum := 0
	for i := 0; i < m.Dimension; i++ {
		sum += m.At(m.Dimension-i-1, i)
	}
	return sum
}

func (m *matrix) AbsoluteDiagonalDifference() int {
	diff := m.Diagonal() - m.ReverseDiagonal()
	if diff < 0 {
		return -diff
	}
	return diff
}
