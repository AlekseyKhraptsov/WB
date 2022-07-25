package main

import "fmt"

func main() {
	mass := []string{"cat", "cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	//Записть уникальных значений в мапу
	for _, v := range mass {
		set[v] = true
	}

	//Вывод результата
	for key := range set {
		fmt.Println(key)
	}
}
