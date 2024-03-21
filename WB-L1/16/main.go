package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	// Разделяем массив на элементы меньше, равные и больше опорного
	var left, right []int
	for i, v := range arr {
		if i == pivotIndex {
			continue
		}
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	sorted := append(append(quicksort(left), pivot), quicksort(right)...)

	return sorted
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	sortedArr := quicksort(arr)

	fmt.Println("Отсортированный массив:", sortedArr)
}

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
