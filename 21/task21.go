package main

import "fmt"

// Интерфейс работает с файлами Xml
type Xlm interface {
	SendXmlData()
}

// Структура принимает файлы с интерфейсом Xml
type XmlDoc struct {
}

func (doc XmlDoc) SendXmlData() {
	fmt.Println("Xml")
}
