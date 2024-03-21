package main

import (
	"fmt"
)

func main() {

	var data [5]int = [5]int{2, 4, 6, 8, 10}

	for _, value := range data {
		go fmt.Println(value * value)
	}
	fmt.Scanln()
}

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
