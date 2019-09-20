package main

import (
	"fmt"
)

var wantp bool = false
var wantq bool = false

func p() {
	for {
		fmt.Println("P NCS1")
		fmt.Println("P NCS2")

		for wantq {
		}
		wantp = true
		fmt.Println("P Critical Section 1")
		fmt.Println("P Critical Section 2")
		wantp = false
	}
}
func q() {
	for {
		fmt.Println("Q NCS1")
		fmt.Println("Q NCS2")

		for wantp {
		}
		wantq = true
		fmt.Println("Q Critical Section 1")
		fmt.Println("Q Critical Section 2")
		wantq = false
	}
}

func main() {
	go p()
	q()
}
