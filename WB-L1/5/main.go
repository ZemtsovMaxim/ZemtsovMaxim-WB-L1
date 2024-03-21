package main

import (
	"context"
	"fmt"
	"time"
)

const n = 5

func main() {

	ch := make(chan int)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*n)

	// Записываем моковые данные в канал
	go func() {
		for {
			ch <- time.Now().Second()
			time.Sleep(time.Second)
		}
	}()

	//Получаем данные до истечения времени
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Конец")
			return
		case <-ch:
			fmt.Printf("Ваши данные: %d\n", <-ch)
		}
	}
}

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/
