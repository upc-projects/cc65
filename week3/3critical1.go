package main

import (
	"fmt"
)

var turn int = 1

func p() {
	for {
		fmt.Println("P NCS1")
		fmt.Println("P NCS2")

		for turn != 1 {
		}

		fmt.Println("P Critical Section 1")
		fmt.Println("P Critical Section 2")

		turn = 2
	}
}
func q() {
	for {
		fmt.Println("Q NCS1")
		fmt.Println("Q NCS2")

		for turn != 2 {
		}

		fmt.Println("Q Critical Section 1")
		fmt.Println("Q Critical Section 2")

		turn = 1
	}
}

func main() {
	go p()
	q()
}
