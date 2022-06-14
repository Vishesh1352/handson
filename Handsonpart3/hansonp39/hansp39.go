package main

import "fmt"

func main() {
	i := "hello"
	switch i {
	case "bye":
		fmt.Println("i is goining")
	case "hello":
		fmt.Println("i is telling hello")
	default:
		fmt.Println("i is speaking gibrish")
	}
}
