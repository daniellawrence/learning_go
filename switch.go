package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("ITs the weekend")
	default:
		fmt.Println("it's a weekday")

	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("ITs before noon", "..", t.Hour(), "am");
	default:
		fmt.Println("ITs after noon", "..", t.Hour()-12, "pm");
	}
}
