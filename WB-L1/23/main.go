package main

import "fmt"

func deleteFromSlice(data []int, num int) []int {

	mapa := make(map[int]int)
	result := []int{}

	for i, _ := range data {
		mapa[i] = data[i]
	}

	delete(mapa, num)

	fmt.Println(mapa)

	for i := 0; i <= len(mapa); i++ {
		if mapa[i] == 0 {
			continue
		} else {
			result = append(result, mapa[i])
		}
	}

	return result
}

func main() {

	data := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(deleteFromSlice(data, 1))

}

//Удалить i-ый элемент из слайса.
