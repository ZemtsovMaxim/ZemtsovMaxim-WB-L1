package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	<-time.After(t)
}

func main() {

	sleep(time.Second * 2)
	fmt.Println("Hello, Wb")
}

//Реализовать собственную функцию sleep.
