package main

import "fmt"

func findType(data interface{}) {
	switch data.(type) {
	default:
		fmt.Printf("unexpected type %T", data) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", data)
	case int:
		fmt.Printf("integer %d\n", data)
	case string:
		fmt.Printf("string %t\n", data)
	case chan int:
		fmt.Printf("chan int %d\n", data)
	case chan string:
		fmt.Printf("chan string %d\n", data)
	case chan bool:
		fmt.Printf("chan bool %d\n", data)
	}

}

func main() {

	var value interface{} = 123

	findType(value)

	var value2 interface{} = "Привет, Wb"

	findType(value2)

	var ch chan bool

	findType(ch)

}

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
