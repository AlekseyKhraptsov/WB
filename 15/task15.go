package main

import (
	"fmt"
	"log"
	"math/rand"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	//1. Поскольку в исходном коде указывалась только правая граница слайса,
	//то в случае, если символов строки меньше 100, вероятна паника, которую
	//допускать нельзя
	rightIndex := 100
	if rightIndex > len(v)-1 {
		log.Println("someFunc(): выход за пределы слайса")
	} else {
		justString = v[:100]
	}
}

// Реализация функции создания строки, посколько ее отсутствие в исходнике
// Вызывало ошибку на этапе компиляции
func createHugeString(length int) string {
	result := make([]byte, length)
	for n := range result {
		result[n] = byte(rand.Intn(122-48) + 48)
	}

	return string(result)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
