package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

// Structure definition
type exampleStruct struct {
	// Members definition
	// --> Non exported members/variables start with a lowercase
	channel chan int
	sum     int
	// --> Exported members/vars start with an upper case
	Exported string `json:"exported_var,omitempty"`
}

// main func is the entrypoint of the main package
func main() {

	// Create a new example struct
	s := new(exampleStruct)
	// Initialize a channel of int, unbuffered
	s.channel = make(chan int)

	// Goroutine with inline function
	go func(s *exampleStruct) {
		// For loop ~= while true
		for {
			// Select case statement
			select {
			// Listen to channel save received input into nbr
			case nbr := <-s.channel:
				// Add nbr to sum of s
				s.sum += nbr
				// Printf with formatting
				fmt.Printf("Received: [%03d]\n", nbr)
			}
		}
	}(s)

	// For loop
	for i := 0; i <= 10; i++ {
		// Send on channel of s, a pseudo random number, within [0,n)
		s.channel <- rand.Intn(100)
	}
	// Println
	fmt.Println("Total", s.sum)

	// Parse sum into string
	s.Exported = strconv.Itoa(s.sum)

	// JSON marshal s into a slice of byte
	b, err := json.Marshal(s)
	// If an error occurs, log.Fatal, exit
	if err != nil {
		log.Fatalf("marshalling JSON, %s", err.Error())
	}

	// Convert []byte into a string
	fmt.Println(string(b))
}
