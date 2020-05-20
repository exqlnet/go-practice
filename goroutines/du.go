package main

import (
	"os"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	aft := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			println("hello world!")
		case <-aft:
			println("goodbye world.")
			os.Exit(0)
		}
	}

}
