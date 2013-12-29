package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}


func main() {
	a, b  := vals()
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
}
