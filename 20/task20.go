package main

import "fmt"

func main() {
	messageString := "snow dog sun"

	mSlice := make(map[int]string)

	key := 0
	var word string

	for _, i := range messageString {

		if string(i) != " " {
			word += string(i)
		} else {
			mSlice[key] = word
			word = ""
			key++
		}
	}
	mSlice[key] = word // Записываем слово в карту

	var newMessage string

	// Перебираем все значения в карте и прибавляем их к новой переменной
	for i := 0; i < len(mSlice); i++ {
		newMessage = mSlice[i] + " " + newMessage
	}

	fmt.Println(newMessage)
}
