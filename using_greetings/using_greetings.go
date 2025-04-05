package main

import (
	"fmt"

	"golearning/greetings"

	"log"
)

func main() {
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Mateus", "Gladys", "Fellipe"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range messages {
		print(v)
		fmt.Println()
	}
}
