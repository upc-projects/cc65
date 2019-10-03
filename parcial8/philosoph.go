package main

import (
	"fmt"
	"sync"
)

func philosophers(name string, lfork, rfork sync.Mutex) {
	for {
		fmt.Println(name, " pensando")
		lfork.Lock()
		rfork.Lock()
		fmt.Println(name, " comiendo")
		lfork.Unlock()
		rfork.Unlock()
	}
}

func main() {
	n := 4
	fork := make([]sync.Mutex, n)
	names := []string{"andrecito", "gmi2", "miguelito", "chino"}
	for i := 0; i < n-1; i++ {
		go philosophers(names[i], fork[i], fork[i+1])
	}
	philosophers(names[n-1], fork[n-1], fork[0])
}
