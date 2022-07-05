package golang_united_school_homework

import (
	"errors"
)

var errOutForMaxRange = errors.New("index out of max range")
var errNoShape = errors.New("no shape by this index")

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}

	return errors.New("max number of shapes")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i > len(b.shapes)-1 {
		return nil, errOutForMaxRange
	}

	if b.shapes[i] == nil {
		return nil, errNoShape
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i > len(b.shapes)-1 {
		return nil, errOutForMaxRange
	}

	if b.shapes[i] == nil {
		return nil, errNoShape
	}

	var shape = b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i > len(b.shapes)-1 {
		return nil, errors.New("index out of max range")
	}

	if len(b.shapes) < i {
		return nil, errors.New("index out of current range")
	}

	var removedShape = b.shapes[i]

	b.shapes[i] = shape

	return removedShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}

	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var isAnyCircle = false
	for i := 0; i < len(b.shapes); i++ {
		if isCircleType(b.shapes[i]) {
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			isAnyCircle = true
		}
	}

	if !isAnyCircle {
		return errors.New("no circles in box")
	}

	return nil
}

func isCircleType(s Shape) bool {
	switch s.(type) {
	case Circle:
		return true
	default:
		return false
	}
}
