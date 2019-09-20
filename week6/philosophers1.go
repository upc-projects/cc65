package main

import (
    "fmt"
    "sync"
)

func philosopher(name string, leftFork, rightFork sync.Mutex) {
    for {
        fmt.Println(name, " pesando")
        leftFork.Lock()
        rightFork.Lock()
        fmt.Println(name, " comiendo")
        leftFork.Unlock()
        rightFork.Unlock()
    }
}

func main() {
    n := 5
    fork := make([]sync.Mutex, n)
    names := []string{ "Socrates", "Platon", "Descartes", "Aristoteles", "nietzsche" }
    for i := 0; i < n - 1; i++ {
        go philosopher(names[i], fork[i], fork[i + 1])
    }
    philosopher(names[n - 1], fork[n - 1], fork[0])
}
