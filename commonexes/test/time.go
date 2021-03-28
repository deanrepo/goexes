package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	tS := "2019-12-30T16:50:50Z"
	t, err := time.Parse(time.RFC3339, tS)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
}
