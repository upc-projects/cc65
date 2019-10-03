package main

import (
	"fmt"
	"sync"
)

var total int

func goingOut(canEnter, canLeave *sync.Mutex, n int) {
	for {
		canLeave.Lock()
		if total != 0 {
			fmt.Println("Client going out")
			total--
			canLeave.Unlock()
		} else {
			canEnter.Unlock()
		}
	}
}

func goingIn(canEnter, canLeave *sync.Mutex, n int) {

	for {
		canEnter.Lock()
		if total < n {
			total++
			fmt.Println("Client entering")
			canEnter.Unlock()
		} else {
			canLeave.Unlock()
		}
	}

}

func main() {

	n := 5
	canEnter := &sync.Mutex{}
	canLeave := &sync.Mutex{}

	canLeave.Lock()

	go goingOut(canEnter, canLeave, n)
	goingIn(canEnter, canLeave, n)

}
