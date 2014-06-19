package main

import "fmt"

func main() {

	msg := ""
	for i := 1; i <= 100; i++ {
		msg = ""
		if i%3 == 0 {
			msg += "fizz"
		}
		if i%5 == 0 {
			msg += "buzz"
		}
		if len(msg) > 0 {
			fmt.Println(msg)
		} else {
			fmt.Println(i)
		}
			
	}
}
