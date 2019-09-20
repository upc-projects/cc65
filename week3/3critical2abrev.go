package main

var wantp bool = false
var wantq bool = false

func p() {
	for {
		for wantq {
		}
		wantp = true
		wantp = false
	}
}
func q() {
	for {
		for wantp {
		}
		wantq = true
		wantq = false
	}
}

func main() {
	go p()
	q()
}
