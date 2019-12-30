package factory

import "fmt"

// Circle contains gthe radius needed to perform operations
type Circle struct {
	radius int
}

func newCircle() *Circle {
	return &Circle{}
}

// Draw draws the circle
func (c Circle) Draw() {
	fmt.Println("Drawing circle...")
}
