package main

import "fmt"

type JsonDoc struct {
}

func (doc JsonDoc) ConvertToXml() string {
	return "xml"
}

type JsonDocAdapter struct {
	jsonDoc *JsonDoc
}

// данный адаптер позволяет конвертировать и отправлять данные из JSON в Xlm
func (adapter JsonDocAdapter) SendXmlData() {
	adapter.jsonDoc.ConvertToXml()
	fmt.Println("adapter XML")
}
