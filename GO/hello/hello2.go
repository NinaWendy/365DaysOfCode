package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Nina")
	if err != nil{
		log.Fatal(err)
	}
	//message := greetings.Hello("Nina")
	fmt.Println(message)
}
