//Остановка горутины с помощью канала

package main

import (
	"fmt"
	"time"
)

func someFunction() int {
	return 1
}

func main() {

	ch := make(chan int, 100)
	done := make(chan bool)
	go func() {
		for {
			select {
			case ch <- someFunction():
			case <-done:
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for i := range ch {
		fmt.Println("receive value: ", i)
	}

	fmt.Println("finish")
}
