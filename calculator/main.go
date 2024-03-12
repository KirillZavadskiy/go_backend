package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Мини калькулятор")
	w.Resize(fyne.NewSize(300, 400))

	label1 := widget.NewLabel("Введите первое число:")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Введите знак +, -, *, /")
	entry2 := widget.NewEntry()

	label3 := widget.NewLabel("Введите второе число:")
	entry3 := widget.NewEntry()

	answer := widget.NewLabel("")
	button := widget.NewButton("Вычислить", func() {
		n1, err := strconv.ParseFloat(entry1.Text, 64)
		n3, er := strconv.ParseFloat(entry3.Text, 64)

		if err != nil || er != nil {
			answer.SetText("Ошибка ввода числа.")
		} else if entry2.Text == "+" {
			sum := n1 + n3
			answer.SetText(fmt.Sprintf("Сумма = %f", sum))
		} else if entry2.Text == "-" {
			sub := n1 - n3
			answer.SetText(fmt.Sprintf("Разница  = %f", sub))
		} else if entry2.Text == "*" {
			mul := n1 * n3
			answer.SetText(fmt.Sprintf("Умножение = %f", mul))
		} else if entry2.Text == "/" {
			div := n1 / n3
			answer.SetText(fmt.Sprintf("Деление = %f", div))
		} else {
			answer.SetText("Ошибка ввода знака.")
		}
	})

	w.SetContent(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		label3,
		entry3,
		button,
		answer,
	))
	w.ShowAndRun()
}