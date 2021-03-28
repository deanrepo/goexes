package main

import "fmt"

func main() {
	myChan := make(chan int)
	select {
	case <-myChan:
	default:
		fmt.Println("Not found")
	}
}
