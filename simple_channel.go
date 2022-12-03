package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Running channels mode...")
	c := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			log.Println("Sending", i)
			c <- i
			time.Sleep(1 * time.Second)
		}

		defer close(c)
	}()

	for v := range c {
		log.Println("Recieved", v)
	}

	time.Sleep(15 * time.Second)
}
