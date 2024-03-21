package main

import (
	"fmt"
)

func binarySearch(arr []int, num int) (int, bool) {

	var mid, result int

	min := 0
	max := len(arr) - 1

	for min <= max {
		mid = (min + max) / 2
		result = mid
		if result == num {
			return result, true
		}
		if result > num {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return 0, false
}

func main() {
	//Массив должен быть отсортированным
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(binarySearch(arr, 3))

}

//Реализовать бинарный поиск встроенными методами языка.
