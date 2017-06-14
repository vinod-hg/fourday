package main

import "fmt"
import "time"

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)
}

func sayHello() {
	fmt.Println("hello")
}
