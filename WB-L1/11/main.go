package main

import "fmt"

// Функция внутри которой создается мапа и заполняется данными из первого слайса - true(ключ - значение).
// Потом пробегая по значениям из второго слайса мы добавляем в результат те значения которые уже были в первом и сейчас есть во втором слайсах.
func findingIntersection(first []int, second []int) []int {

	mapa := make(map[int]bool)
	result := []int{}

	for _, v := range first {
		mapa[v] = true
	}

	for _, v := range second {
		if mapa[v] != false {
			result = append(result, v)
		}
	}

	return result
}

func main() {

	firstSet := []int{1, 2, 55, 3, 4, 5, 6, 8, 11}
	secondSet := []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 55}

	fmt.Println(findingIntersection(firstSet, secondSet))

}

//Реализовать пересечение двух неупорядоченных множеств.
