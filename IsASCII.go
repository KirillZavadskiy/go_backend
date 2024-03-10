package main

import (
	"fmt"
)

func IsASCII(s string) bool {
	res := true
	for _, lat := range s {
		if rune(lat) > 0 && rune(lat) < 255 {
			res = true
		} else {
			res = false
			return res
		}
	}

	return res
}

func main() {
	f := "Привет world"
	t := "world"
	fmt.Println("Должно быть false -", IsASCII(f))
	fmt.Println("Должно быть true -", IsASCII(t))
}
