package main

import "fmt"

func main() {
	v1 := struct {
		doors int
		color string
	}{
		doors: 4,
		color: "black",
	}
	fmt.Println(v1)
}
