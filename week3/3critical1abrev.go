package main

var turn int = 1

func p() {
	for {
		for turn != 1 {
		}
		turn = 2
	}
}
func q() {
	for {
		for turn != 2 {
		}
		turn = 1
	}
}

func main() {
	// Programa no hace nada, si hubiera deadlock, GO lo detectaría
	go p()
	q()
}
