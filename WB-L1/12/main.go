package main

import "fmt"

// Функция в которой мапа заполнсяется данными из исходного слайса(заодно посчитаем сколько раз повторяется то или иное слово),
// потом из мапы заполняем слайс с результатами по индексу.
func convertingToSet(slice []string) []string {

	mapa := make(map[string]int)
	result := []string{}
	for _, v := range slice {
		mapa[v] = mapa[v] + 1
		fmt.Println(mapa)
	}

	for i, _ := range mapa {
		result = append(result, i)
		fmt.Println(mapa[i])

	}

	return result
}

func main() {

	data := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(convertingToSet(data))

}

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
