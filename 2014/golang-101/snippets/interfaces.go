package main

import "fmt"
import "math"

// START OMIT
type geometry interface {
	area() float64
}

type square struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Printf("area, %v\n", g.area())
}

// END OMIT

// START OMITu // OMIT
func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}
	measure(s)
	measure(c)
}

// END OMITu // OMIT
