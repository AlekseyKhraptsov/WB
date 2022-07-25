package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func intersect(a, b []int) []int {
	aSet := make(map[int]struct{})
	bSet := make(map[int]struct{})

	for _, x := range a {
		aSet[x] = struct{}{}
	}
	for _, x := range b {
		bSet[x] = struct{}{}
	}
	//Проверяем, есть ли ключ из первого во втором и если есть - выписываем его в out []int
	var out []int
	for x := range aSet {
		_, ok := bSet[x]
		if ok {
			out = append(out, x)
		}
	}
	return out
}

func main() {
	fmt.Println(intersect([]int{1, 2, 3, 4}, []int{1, 2, 5, 6, 1}))
}
