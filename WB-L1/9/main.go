package main

import (
	"fmt"
	"sync"
)

func main() {

	ch1 := make(chan int, 10)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	var data [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// создаем в цикле горутину в которой пишем в ch1 значения от x до n (в данном случае от 0 до 10), а в ch2 их квадраты
	for i, _ := range data {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch1 <- data[i]
			ch2 <- i * 2
		}(i)
		//выводим данные из второго канала
		fmt.Println(<-ch2)
	}
	wg.Wait()
}

/*
Разработать конвейер чисел.
 Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/
