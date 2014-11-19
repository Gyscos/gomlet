package gomlet

import (
	"errors"
	"log"
)

type Matrix struct {
	Values []float64
	Width  int
	Height int
}

func NewMatrix(w, h int) Matrix {
	return Matrix{make([]float64, h*w), w, h}
}

func (m Matrix) Get(x, y int) float64 {
	return m.Values[x+m.Width*y]
}

func (m Matrix) Set(x, y int, value float64) {
	m.Values[x+m.Width*y] = value
}

func MatTimes(A, B Matrix) (Matrix, error) {
	result := NewMatrix(B.Width, A.Height)
	if A.Width != B.Height {
		return result, errors.New("Matrix sizes do not match")
	}

	for x := 0; x < result.Width; x++ {
		for y := 0; y < result.Height; y++ {
			sum := 0.0
			for i := 0; i < A.Width; i++ {
				sum += A.Get(i, y) * B.Get(x, i)
			}
			result.Set(x, y, sum)
		}
	}

	return result, nil
}

func (m Matrix) Apply(v []float64) []float64 {
	A := NewMatrix(len(v), 1)
	A.Values = v
	B, err := MatTimes(A, m)
	if err != nil {
		log.Fatal(err)
	}
	return B.Values
}
