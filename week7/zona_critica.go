// lectores y escritores. un solo escritor puede entrar, varios lectores
// pueden entrar a la vez

package main

import (
	"fmt"
	"sync"
)

var counter = 0

func writer(roomEmpty *sync.Mutex) {
	for {
		roomEmpty.Lock()
		fmt.Println("writting 1")
		fmt.Println("writting 2")
		fmt.Println("writting 3")
		roomEmpty.Unlock()
	}
}

func reader(roomEmpty, mutex *sync.Mutex) {
	for {
		mutex.Lock()
		counter++
		if counter == 1 {
			roomEmpty.Lock()
		}
		mutex.Unlock()

		fmt.Println("reading 1")
		fmt.Println("reading 2")
		fmt.Println("reading 3")

		mutex.Lock()
		counter--

		if counter == 0 {
			roomEmpty.Unlock()
		}
		mutex.Unlock()

	}
}
func main() {
	roomEmpty := &sync.Mutex{}
	mutex := &sync.Mutex{}

	go reader(roomEmpty, mutex)
	writer(roomEmpty)
	go reader(roomEmpty, mutex)

}
