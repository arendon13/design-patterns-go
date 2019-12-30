package factory

// ShapeType is an enum of type int used to determine shape type
type ShapeType int

const (
	rectangleType ShapeType = 1 << iota
	squareType    ShapeType = 1 << iota
	circleType    ShapeType = 1 << iota
)

// NewShape gets the shape we will be working with
func NewShape(shapeType ShapeType) Shape {

	switch shapeType {
	case rectangleType:
		return newRectangle()
	case squareType:
		return newSquare()
	case circleType:
		return newCircle()
	default:
		return nil
	}

}
