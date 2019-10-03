package main

import "fmt"

var sum int

func suma(a int, b int) {
	sum = a + b
	fmt.Println("suma ", sum)
}

func sumSlices(slices []int) {
	for _, num := range slices {
		go suma(num, sum)
	}
}

func main() {
	var slice = []int{12, 43, 2, 324, 23, 121}
	sumSlices(slice)
}
