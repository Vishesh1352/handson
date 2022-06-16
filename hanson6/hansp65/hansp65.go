// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}
type squar struct {
	l float64
	w float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (s squar) area() float64 {
	return s.l * s.w
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {
	cir := circle{
		r: 1.23,
	}
	sqr :=
		squar{
			l: 5.2,
			w: 5.2,
		}
	fmt.Println("arear of circle\n")
	info(cir)
	fmt.Println("area of squar\n")
	info(sqr)
}
