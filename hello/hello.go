package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // remove date and time from the log message

	message, err := greetings.Hello("Pooooob")
	if err != nil {
		log.Fatal(err) // this call also exits the program
	}

	fmt.Println(">>>>>>>", message)
}
