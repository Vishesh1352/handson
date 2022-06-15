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

	fmt.Println(p1)
	m := map[string]person{
		p1.last: p1,
	}

	fmt.Println(m)

	for k, v := range p1.icecreame_flavor {
		fmt.Println(k, v)
	}

}
