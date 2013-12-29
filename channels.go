package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {messages <- "ping 1"}()

	msg := <- messages
	fmt.Println(msg)
	
}
