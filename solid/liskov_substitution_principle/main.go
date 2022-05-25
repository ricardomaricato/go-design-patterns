package main

import "fmt"

type Area interface {
	GetArea() int
}

type IRectangle interface {
	Area
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

//     vvv !! POINTER

func (r *Rectangle) GetArea() int {
	return r.width * r.height
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// modified LSP
// If a function takes an interface and
// works with a type T that implements this
// interface, any structure that aggregates T
// should also be usable in that function.

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

type ISquare interface {
	Area
	SetSize(size int)
	GetSize() int
}

type Square2 struct {
	size int
}

func (s *Square2) GetArea() int {
	return s.size * s.size
}

func (s *Square2) SetSize(size int) {
	s.size = size
}

func (s *Square2) GetSize() int {
	return s.size
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}
func SquareUseIt(sized ISquare) {
	sized.SetSize(10)
	width := sized.GetSize()
	expectedArea := 10 * width
	actualArea := sized.GetArea()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}

func RetangleUseIt(sized IRectangle) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetArea()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}

func main() {
	rc := &Rectangle{2, 3}
	RetangleUseIt(rc)

	sq := NewSquare(5)
	RetangleUseIt(sq)

	sq2 := &Square2{30}
	SquareUseIt(sq2)
}
