package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	z = x*x
	return z
}

func main() {
     fmt.Printf("Hello World! - %v\n", Sqrt(123))

}