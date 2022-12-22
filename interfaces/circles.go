package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Square struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Height: %.2f, Width: %.2f}", s.Height, s.Width)
}

type Sizer interface {
	Area() float64
}

type Shaper interface {
	Sizer
	fmt.Stringer
}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}

	return s2
}

func PrintArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}
