package main

import (
	"fmt"
	"study/hello"
)

func main() {
	fmt.Println("Welcome to Study with Go...")

	fmt.Println(hello.Greeting())

	greeting, error := hello.SayHelloTo("")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(greeting)
	}
	
}