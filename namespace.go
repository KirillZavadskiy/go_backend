package main

import (
	"fmt"
)

func GenerateSelfStory(name string, age int, money float64) string {
	r := fmt.Sprintf(
		"Привет! Меня зовут %s. Мне %d лет. Я изучаю Python и Go более %.1f года.", name, age, money,
	)
	return r
}

func main() {
	fmt.Print(GenerateSelfStory("Кирилл", 30, 1.0))
}
