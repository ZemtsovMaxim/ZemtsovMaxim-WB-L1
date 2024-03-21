package main

import (
	"fmt"
	"math/big"
)

func add(a, b *big.Float) *big.Float {
	return new(big.Float).Add(a, b)
}
func subtract(a, b *big.Float) *big.Float {
	return new(big.Float).Sub(a, b)
}
func divide(a, b *big.Float) *big.Float {
	return new(big.Float).Quo(a, b)
}
func multiply(a, b *big.Float) *big.Float {
	return new(big.Float).Mul(a, b)
}

func main() {

	a, _ := new(big.Float).SetPrec(512).SetString("1048577")
	b, _ := new(big.Float).SetPrec(512).SetString("1048577")
	fmt.Println(a, b)
	fmt.Println(add(a, b))
	fmt.Println(subtract(a, b))
	fmt.Println(divide(a, b))
	fmt.Println(multiply(a, b))

}

//	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
