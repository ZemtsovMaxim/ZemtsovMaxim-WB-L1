package main

import (
	"fmt"
	"time"
)

// Функция для чтения данных из канала
func readFromCh(name int, readerChan <-chan int) {
	for {
		timeSec := <-readerChan
		fmt.Printf("Воркер %d получил данные: %d\n", name, timeSec)
	}
}

func main() {
	var workers int
	fmt.Println("Введите количество воркеров: ")
	fmt.Scanf("%d\n", &workers)

	// Создаем канал
	ch := make(chan int)

	for i := 1; i <= workers; i++ {
		go readFromCh(i, ch)
	}

	// Записываем моковые данные в канал
	for {
		ch <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
