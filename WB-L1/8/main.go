package main

import (
	"fmt"
)

func setBit(value int64, position int, bitValue int) int64 {
	// Создаем маску с установленным битом в позиции position
	mask := int64(1) << position

	if bitValue == 1 {
		value |= mask
	} else {
		value &= ^mask
	}

	return value
}

func main() {

	var num int64 = 16
	position := 2
	bitValue := 1

	fmt.Printf("Исходное значение: %d\n", num)

	num = setBit(num, position, bitValue)

	fmt.Printf("Значение после установки %d-го бита в %d: %d\n", position, bitValue, num)
}

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
