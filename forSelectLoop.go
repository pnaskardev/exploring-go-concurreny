package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {

	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
}

func ForLoop() {

	// PRIMITIVES
	// Primitives()

	// For Select Loop
	// ForSelectLoop()

	charChannel := make(chan string, 3)

	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		select {
		case charChannel <- char:
		}
	}
	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

	doneChannel := make(chan bool)

	go doWork(doneChannel)

	time.Sleep(time.Second * 2)

	close(doneChannel)
}
