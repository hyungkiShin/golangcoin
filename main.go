package main

import "fmt"

type person struct {
	contry string
	age    int
}

func (p person) sayHello() {
	fmt.Printf("Hello My contry is %s and Korean Age is %d", p.contry, p.age)
}

func main() {
	nico := person{contry: "korean", age: 29}
	nico.sayHello()
}
