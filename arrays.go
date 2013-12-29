package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[3] = 100
	fmt.Println(a[4])
	fmt.Println("a:", a)

	fmt.Println(len(a))

	var x[4][4]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(i,j)
			x[i][j] = i
		}
	}

	fmt.Println(x)

}
