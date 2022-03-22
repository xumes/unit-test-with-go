package area

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

func CalculateArea(s Shape) {
	fmt.Printf("The Area of this shape is %0.2f", s.Area())
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
