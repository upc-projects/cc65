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

		wantp = true
		for wantq {
		}
		fmt.Println("P Critical Section 1")
		fmt.Println("P Critical Section 2")
		wantp = false
	}
}
func q() {
	for {
		fmt.Println("Q NCS1")
		fmt.Println("Q NCS2")

		wantq = true
		for wantp {
		}
		fmt.Println("Q Critical Section 1")
		fmt.Println("Q Critical Section 2")
		wantq = false
	}
}

func main() {
	go p()
	q()
}
