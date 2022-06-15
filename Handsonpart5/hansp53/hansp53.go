package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type seden struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "greay",
		},
		fourWheel: true,
	}
	s1 := seden{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println(t1.vehicle)
	fmt.Println(s1.luxury)

}
