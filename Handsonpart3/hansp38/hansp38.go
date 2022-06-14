package main

import "fmt"

func main() {
	i := 100
	switch {
	case i == 50:
		fmt.Println("i is 50")
	case i < 50:
		fmt.Println("i is less than 50")
	case i < 200:
		fmt.Println("i is less than 200")
	}
}
