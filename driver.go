package main

import (
	"github.com/arendon13/design-patterns-go/factory"
)

func main() {

	runFactory()

}

func runFactory() {

	someShape := factory.NewShape(1)
	someShape.Draw()

	someShape = factory.NewShape(2)
	someShape.Draw()

	someShape = factory.NewShape(4)
	someShape.Draw()

}
