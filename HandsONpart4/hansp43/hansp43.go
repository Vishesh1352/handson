package main

import "fmt"

func main() {
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(b[0:5])
	fmt.Println(b[5:10])
	fmt.Println(b[2:7])
	fmt.Println(b[1:6])
}
