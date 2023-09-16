package main

import (
	"fmt"
	"log"
)

func main2() {
	log.Println("Start")

	message := buildMessage()
	fmt.Println(message)

	log.Println("Finish")
}
