package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {

	fmt.Println(num)
}

func goRoutines() {
	// GO ROUTINE
	// NO PARTICULAR ORDER
	go someFunc("25")

	go someFunc("69")

	// WE ARE RESPONSIBLE TO MAKE THE THREAD JOIN AGAIN

	// MAKES MAIN SLEEP
	time.Sleep(time.Second * 2)

	fmt.Println("HELLO WORLD")
}

func main() {

	// GROUTINES
	goRoutines()

	// CHANNELS
	

}
