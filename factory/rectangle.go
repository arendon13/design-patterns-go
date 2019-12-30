package factory

import "fmt"

// Rectangle contains length and width needed to perform operations
type Rectangle struct {
	length int
	width  int
}

func newRectangle() *Rectangle {
	return &Rectangle{}
}

// Draw draws the rectangle
func (r Rectangle) Draw() {
	fmt.Println("Drawing rectangle...")
}
