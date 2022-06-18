package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func (p *person) speak() {
	fmt.Println(p)
}

type human interface {
	speak()
}

func SaySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		Name: "yash",
		Age:  55,
	}
	SaySomething(&p1)
}
