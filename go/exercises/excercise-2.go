package exercises

import (
	"fmt"
)

type Square struct {
	len   int
	width int
}
type Calculator interface {
	Area() int
}

func (s Square) Area() int {
	return s.len * s.width
}

func Exercise2() {

	square := Square{2, 2}
	var newSquare Calculator = square
	newSquareVal, ok := newSquare.(Square)
	fmt.Println("Exercise2:  ", newSquareVal, ok)
}
