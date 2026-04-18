package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go slowGreet("Nice to meet you!", done)
	go greet("I'm learning Go!", done)
	go slowGreet("I'm learning Go!", done)

	for doneChan := range done {
		fmt.Println(doneChan)
	}
}
