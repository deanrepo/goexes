package main

import (
	"fmt"
	"time"
)

func portal1(channel1 chan string) {
	time.Sleep(3 * time.Second)
	channel1 <- "Welcome to channel 1"
}

func portal2(channel2 chan string) {
	time.Sleep(9 * time.Second)
	channel2 <- "Welcome to channel 1"
}

func main() {
	r1 := make(chan string)
	r2 := make(chan string)

	go portal1(r1)
	go portal2(r2)

	select {
	case op1 := <-r1:
		fmt.Println(op1)
	case op2 := <-r2:
		fmt.Println(op2)
	}
}
