// Остановка горутины с помощью WaitGroup
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for {
			value, ok := <-ch
			if !ok {
				wg.Done()
				return
			}
			fmt.Println(value)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {

		ch <- 10

	}
	close(ch)
	wg.Wait()
	fmt.Println("Конец")
}
