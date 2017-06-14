package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		mu.Lock()
		m["1"] = "a" // First conflicting access.
		mu.Unlock()
		c <- true
	}()
	mu.Lock()
	m["2"] = "b" // Second conflicting access.
	mu.Unlock()
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
