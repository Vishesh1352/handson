package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {

	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	var sq1 = sqrtError{
		lat:  "yash",
		long: "hour",
		err:  fmt.Errorf("this is sq1 error"),
	}
	if f < 0 {
		// write your error code here
		fmt.Printf("%v", sq1)
	}
	return 42, nil
}
