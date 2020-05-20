package main

import (
	"fmt"
	"sync"
)

var stat = 1
var mu = sync.RWMutex{}
var stats = []string{"\\", "|", "-", "/", "-"}

func Stat() string {
	mu.Lock()
	defer mu.Unlock()
	stat++
	if stat >= len(stats) {
		stat = 0
	}
	return stats[stat]
}

func SetStat(s string)  {
	mu.Lock()
	defer mu.Unlock()
}

func main() {
	for i := 0; i < 1000000; i++ {
		go func() {
			for {
				fmt.Printf("\r%s", Stat())
			}
		}()
	}
	c := make(chan struct{})
	<- c
}
