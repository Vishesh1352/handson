package main

import "fmt"

type person struct {
	first            string
	last             string
	icecreame_flavor []string
}

func main() {
	p1 := person{
		first:            "james",
		last:             "jolly",
		icecreame_flavor: []string{"vanila"},
	}
	p2 := person{
		first:            "johny",
		last:             "english",
		icecreame_flavor: []string{"choclate"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.icecreame_flavor)
	for i, v := range p1.icecreame_flavor {
		fmt.Println(i, v)
	}
	for i, v := range p2.icecreame_flavor {
		fmt.Println(i, v)
	}
}
