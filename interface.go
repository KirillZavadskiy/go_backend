package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}
type NumbersInterface interface {
	plus() int
	umnozh() int
}

func (n Numbers) plus() int {
	return n.num1 + n.num2
}
func (n Numbers) umnozh() int {
	return n.num1 * n.num2
}
func (n Numbers) minus() int {
	return n.num1 - n.num2
}

func main() {
	var i NumbersInterface
	numbers := Numbers{2, 3}
	i = numbers
	fmt.Printf("Реализация через интерфейс = %d\n", i.plus())
	fmt.Printf("Реализация через интерфейс = %d\n", i.umnozh())
	fmt.Printf("Реализация через метод = %d", numbers.minus())
}
