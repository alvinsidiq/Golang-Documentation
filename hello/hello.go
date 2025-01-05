package main

import (
    "fmt"
	"log"
    "example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

    // A slice of names. 
	names :=[]string{"Gladys","samantha","Darin"}

    // Get a greeting message and print it.
    message,err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}


    fmt.Println(message)
}
