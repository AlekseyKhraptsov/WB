package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Структура
type Human struct {
	Name string
	Race string
}

// Метод
func (h Human) PrintInfo() {
	fmt.Println("Name:", h.Name)
	fmt.Println("Race:", h.Race)
}

// Делаем наследование
type Action struct {
	Human
}

func main() {
	h := Human{
		Name: "Dakir",
		Race: "Asian",
	}
	a := Action{
		Human: h,
	}
	h.PrintInfo()
	a.PrintInfo()
	a.Human.PrintInfo()
	fmt.Println(a.Race)
}
