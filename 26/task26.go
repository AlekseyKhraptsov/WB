package main

import (
	"fmt"
	"strings"
)

// приведем все символы в нижний регистр и проверим на уникальность
func unique(s string) bool {
	repeats := make(map[rune]bool)
	for _, a := range strings.ToLower(s) {

		if repeats[a] {
			return false
		}
		repeats[a] = true
	}
	return true
}

func main() {
	slice := []string{"abcd", "abCdefAaf", "aAbcd"}
	for _, s := range slice {
		fmt.Printf("%s %t\n", s, unique(s))
	}
}
