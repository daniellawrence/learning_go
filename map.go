package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 14

	fmt.Println(m)

	v1 := m["k1"]

	fmt.Println(v1)
	fmt.Println(m["k1"])

	x, y := m["k2"]
	fmt.Println(x, y)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

}
