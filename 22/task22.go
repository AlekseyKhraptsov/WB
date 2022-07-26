package main

import (
	"fmt"
	"math/big"
)

// Задание 22
// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	a := big.NewInt(2 << 22)
	b := big.NewInt(2 << 43)
	fmt.Println(Mul(a, b))
	fmt.Println(Add(a, b))
	fmt.Println(Sub(a, b))
	fmt.Println(Div(a, b))
}

// перемножаем
func Mul(a, b *big.Int) *big.Int {

	return new(big.Int).Mul(a, b)

}

// складываем
func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// вычитаем
func Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// делим
func Div(a, b *big.Int) *big.Float {
	a1 := new(big.Float).SetInt(a)
	b1 := new(big.Float).SetInt(b)
	return new(big.Float).Quo(a1, b1)
}
