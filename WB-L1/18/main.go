package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	c.counter++
	c.Unlock()
}

func main() {
	var c Counter
	var wg sync.WaitGroup

	for i := 0; i < 123; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()
	fmt.Println(c)
}

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.
