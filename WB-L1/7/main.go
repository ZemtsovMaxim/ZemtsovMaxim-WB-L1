package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var p map[int]int
	var wg sync.WaitGroup
	var mu sync.Mutex

	p = make(map[int]int)

	//Создаем n горутин (сейчас 100) в которых мапа заполняется рандомными int(ами)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i2 := i
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			p[i2] = rand.Int()
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	// Выводим map после завершения всех горутин
	fmt.Println(p)
}

//Реализовать конкурентную запись данных в map
