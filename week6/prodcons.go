package main

import (
	"fmt"
)

const n int = 5

func wait(s int) {
	if s > 0 {
		s--
	}
}

func signal(s int) {
	s++
}

func producer(buffer []int, pos int, sbuff int, notEmpty int, notFull int) {
	var d int
	for {
		d++
		wait(notFull)
		wait(sbuff)
		buffer[pos] = d
		pos++
		signal(sbuff)
		signal(notEmpty)
	}
}

func consumer(buffer []int, pos int, sbuff int, notEmpty int, notFull int) {
	var d int
	for {
		wait(notEmpty)
		wait(sbuff)
		pos--
		d = buffer[pos]
		signal(sbuff)
		signal(notFull)
		fmt.Println("Consumiendo ", d)
	}
}

func main() {
	var sbuff = 1
	var notEmpty int
	var pos int
	var notFull = n
	buffer := make([]int, n)

	producer(buffer, pos, sbuff, notEmpty, notFull)
	// go consumer(buffer)

}
