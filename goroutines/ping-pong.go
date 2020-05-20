package main

import (
	"os"
	"time"
)

func main() {
	c := make(chan struct{})
	count := 0
	aft := time.After(1*time.Second)
	go func() {
		for {
			select {
			case <- c:
				count ++
				println(count)
			case c <- struct{}{}:
			}
		}

	}()
	go func() {
		for {
			select {
			case <- c:
				count ++
				println(count)
			case c <- struct{}{}:
			}
		}

	}()

	select {
	case <- aft:
		os.Exit(0)
	}
}