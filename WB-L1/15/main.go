package main

import "fmt"

var justString string

func createHugeString(size int) string {
	// какая-нибудь логика
	return "HugeString"
}

func someFunc() {
	v := createHugeString(1 << 10)

	// Создаем новый срез из первых 100 байтов из v
	copyBytes := make([]byte, 100)
	copy(copyBytes, v)

	// Преобразуем срез в строку
	justString = string(copyBytes)
}

func main() {
	someFunc()
	fmt.Println(justString)
}

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]

func main() {
  someFunc()
}

 Проблема заключается в том, что создается большая строка v с использованием функции createHugeString,
 а затем в переменную justString записывается только ее первые 100 символов.
 Однако, полная строка v остается в памяти, и Go не сможет освободить выделенную для нее память до завершения программы.
*/
