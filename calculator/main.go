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
	icon, err := fyne.LoadResourceFromPath("icon.png")
	if err != nil {
		fmt.Println("Лого не загрузилось")
	}
	w.SetIcon(icon)

	label1 := widget.NewLabel("Введите первое число:")
	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Первое число")

	label2 := widget.NewLabel("Выберите знак:")
	entry2 := widget.NewRadioGroup([]string{"+", "-", "*", "/"}, func(string) {})
	entry2.Horizontal = true

	label3 := widget.NewLabel("Введите второе число:")
	entry3 := widget.NewEntry()
	entry3.SetPlaceHolder("Второе число")

	answer := widget.NewLabel("")
	button := widget.NewButton("Вычислить", func() {
		n1, err := strconv.ParseFloat(entry1.Text, 64)
		n3, er := strconv.ParseFloat(entry3.Text, 64)

		if err != nil || er != nil {
			answer.SetText("Ошибка ввода числа.")
		} else if entry2.Selected == "+" {
			sum := n1 + n3
			answer.SetText(fmt.Sprintf("Сумма = %f", sum))
		} else if entry2.Selected == "-" {
			sub := n1 - n3
			answer.SetText(fmt.Sprintf("Разница  = %f", sub))
		} else if entry2.Selected == "*" {
			mul := n1 * n3
			answer.SetText(fmt.Sprintf("Умножение = %f", mul))
		} else if entry2.Selected == "/" {
			div := n1 / n3
			answer.SetText(fmt.Sprintf("Деление = %f", div))
		} else {
			answer.SetText("Ошибка! Выберите знак")
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
