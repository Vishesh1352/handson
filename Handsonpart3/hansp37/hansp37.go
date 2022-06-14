package main

import "fmt"

func main() {
	i := 100
	if i < 50 {
		fmt.Printf("i is less then 50")
	} else if i < 200 {
		fmt.Printf("i is less then 200")
	} else {
		fmt.Printf("i is something")
	}

}
