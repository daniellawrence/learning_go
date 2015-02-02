package main

import (
	"fmt"
	"time"
)

func main() {
	var numberList = []int{1, 5, 3, 2 ,4, 9, 8, 1, 2, 4, 5}
	done := make(chan int, 1)

	for _, i := range numberList {
		go func(i int){
			time.Sleep(time.Duration(i) * time.Microsecond * 2)
			done <- i
		}(i)
	}
	for _ = range numberList {
		fmt.Println(<-done)
	}
	
}
