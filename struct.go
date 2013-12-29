package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Bob", age: 20})
	fmt.Println(person{name: "Fred"})

	s := &person{name: "Fred", age: 3}
	fmt.Println(s.name)

	s.age = 51
	fmt.Println(s)

}
