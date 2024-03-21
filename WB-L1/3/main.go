package main

import (
	"fmt"
)

func main() {

	var data [5]int = [5]int{2, 4, 6, 8, 10}

	fmt.Println(calculateSquares(data))
}

func calculateSquares(data [5]int) int {
	if len(data) < 1 {
		return 0
	}

	var sum int
	ch := make(chan int, len(data))

	go func() {
		for i := 0; i < len(data); i++ {
			go func(num int) {

				square := num * num
				ch <- square
			}(data[i])
		}
	}()

	for i := 0; i < len(data); i++ {
		sum += <-ch
	}

	close(ch)
	return sum
}

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
