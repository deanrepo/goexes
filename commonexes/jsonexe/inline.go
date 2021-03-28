package main

import "encoding/json"

import "log"

type Person struct {
	Name string  `json:"name"`
	Age  int     `json:"age"`
	Addr Address `json:"address,inline"`
}

type Address struct {
	Street string `json:"street"`
	Number int    `json:"number"`
}

func main() {
	addr := Address{}
	addr.Street = "test street"
	addr.Number = 123
	p1 := &Person{}
	p1.Addr = addr
	p1.Name = "Lily"
	p1.Age = 20

	ret, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(ret))
}
