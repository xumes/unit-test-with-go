package area

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func CalculateArea(s Shape) {
	fmt.Printf("The area of this shape is %0.2f", s.area())
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
