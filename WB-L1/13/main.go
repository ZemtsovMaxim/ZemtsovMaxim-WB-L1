package main

import "fmt"

func main() {

	a := 1
	b := 2
	// 1 способ
	a, b = b, a
	fmt.Println(a, b)
	// 2 способ
	fmt.Println(swap(a, b))
}

func swap(a, b int) (int, int) {
	return b, a
}

//Поменять местами два числа без создания временной переменной.
