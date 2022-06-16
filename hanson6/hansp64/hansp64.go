package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("person name:", p.first, p.last, "and age:", p.age)
}
func main() {
	pi := person{
		first: "james",
		last:  "bond",
		age:   55,
	}
	pi.speak()
}
