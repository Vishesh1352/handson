package main

import "fmt"

func main() {
	fs := []string{"James", "Bond", "Shaken, not stirred"}
	ss := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(fs)
	fmt.Println(ss)
	xxs := [][]string{fs, ss}
	fmt.Println(xxs)

	for i, vs := range xxs {
		fmt.Println("record: ", i)
		for j, v := range vs {
			fmt.Printf("\t index position : %v \t value : \t %v", j, v)
		}
	}
}
