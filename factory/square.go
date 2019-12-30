package factory

import "fmt"

// Square contains length and width needed to perform operations
type Square struct {
	length int
	width  int
}

func newSquare() *Square {
	return &Square{}
}

// Draw draws the square
func (s Square) Draw() {
	fmt.Println("Drawing square...")
}
