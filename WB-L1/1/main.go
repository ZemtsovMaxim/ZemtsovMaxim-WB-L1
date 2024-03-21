package main

import "fmt"

type Human struct {
	name string
	age  int8
}

func (h Human) printAge() {
	fmt.Println(h.age)
}

func (h Human) printName() {
	fmt.Println(h.name)
}

type Action struct {
	description string
	time        int8
	Human
}

func main() {

	var run = Action{description: "running", time: 5, Human: Human{name: "Igor", age: 25}}
	run.printAge()
	run.Human.printAge()
	run.printName()
	run.Human.printName()

}

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
