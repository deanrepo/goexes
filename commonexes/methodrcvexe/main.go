package main

import (
	"fmt"
)

type TestStruct struct {
}

func (TestStruct) print() {
	fmt.Println("this is a test for method receive")
}

func main() {
	ts := TestStruct{}
	ts.print()
	fmt.Println("this is a test for keyboard")
	t := []int{1, 2, 3, 4, 5}
	for _, v := range t {
		fmt.Println(v)
	}
}
