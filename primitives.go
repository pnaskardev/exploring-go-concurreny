package main

import (
	"fmt"
)

func someFunc(num string, channel chan string) {
	channel <- num
}

func Primitives() {

	// GO ROUTINE
	// NO PARTICULAR ORDER

	channel1 := make(chan string)
	channel2 := make(chan string)
	go someFunc("25", channel1)
	go someFunc("69", channel2)

	// WE ARE RESPONSIBLE TO MAKE THE THREAD JOIN AGAIN

	// block the thread until it gets from atleast one message
	select {
	case msg1 := <-channel1:
		fmt.Println(msg1)
	case msg2 := <-channel2:
		fmt.Println(msg2)
	}

	fmt.Println("HELLO WORLD")

}
