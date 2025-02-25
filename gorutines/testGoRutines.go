package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hello from the goroutine!")
	fmt.Println("Hello from the main goroutine!")
	go func(message string) {
		fmt.Println(message)
	}("Hello from the anonymous goroutine!")
	time.Sleep(4*time.Second)
	fmt.Println("Exiting the program...")
	
}

func say(s string) {
	time.Sleep(1*time.Second)
	fmt.Println(s)
}
