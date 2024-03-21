package main

import (
	"fmt"
	"strings"
)

func reverseWords(data string) []string {

	newData := strings.Split(data, " ")

	for i, j := 0, len(newData)-1; i < j; i, j = i+1, j-1 {
		newData[i], newData[j] = newData[j], newData[i]
	}

	return newData
}

func main() {

	var data = "snow dog sun"
	fmt.Println(reverseWords(data))

}

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
