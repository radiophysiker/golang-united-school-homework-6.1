package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

var (
	errorOutOfCapacity   = errors.New("max capacity reached")
	errorIndexOutOfRange = errors.New("index out of range")
	errorNothingRemoved  = errors.New("nothing removed")
)

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errorOutOfCapacity
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errorIndexOutOfRange
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	shapes := append(b.shapes[0:i], b.shapes[i+1])
	b.shapes = shapes
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	replaceShape, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	b.shapes[i] = shape
	return replaceShape, err
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64
	for i := range b.shapes {
		result += b.shapes[i].CalcPerimeter()
	}
	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64
	for i := range b.shapes {
		result += b.shapes[i].CalcArea()
	}
	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	shapes := make([]Shape, 0, len(b.shapes))
	for i := range b.shapes {
		if _, ok := b.shapes[i].(*Circle); !ok {
			shapes = append(shapes, b.shapes[i])
		}
	}
	if len(b.shapes) == len(shapes) {
		return errorNothingRemoved
	}
	b.shapes = shapes
	return nil

}
