package main

import (
	"fmt"
)

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "vishesh"
}
func main() {
	p1 := person{
		name: "cid dha",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
