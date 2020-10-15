package main

import "fmt"

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	fmt.Stringer
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Square of width %f and length %f", r.Width, r.Length)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}
}
