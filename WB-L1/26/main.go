package main

import (
	"fmt"
	"strings"
)

func isUniq(data string) bool {
	strings.ToLower(data)
	mapa := make(map[rune]bool)
	for _, v := range data {
		_, ok := mapa[v]
		if ok {
			return false
		}
		mapa[v] = true
	}
	return true
}

func main() {

	data := "abcdea"
	fmt.Println(isUniq(data))

}

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false

*/
