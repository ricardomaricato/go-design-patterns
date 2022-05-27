package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	Name string
	Age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hello, I'm %s, I'm have %d years old.\n", p.Name, p.Age)
}

type tiredPerson struct {
	Name string
	Age  int
}

func (t *tiredPerson) SayHello() {
	fmt.Printf("Sorry, I'm too tired")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	// initialize directly
	p := person{"John", 22}
	fmt.Println(p)

	// use a constructor
	p2 := NewPerson("Jane", 21)
	p2.SayHello()

	p3 := NewPerson("Jane", 101)
	p3.SayHello()
}
