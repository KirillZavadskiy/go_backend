// Работа со структурой с использованием указателей
package main

import "fmt"

type User struct {
	name  string
	age   int16
	score []int
}

func changeAge(u *User) {
	u.age = 23
}
func (u *User) setName(name1 string) string {
	u.name = name1
	return u.name
}
func (u User) getScore() int {
	max := 0
	for _, sc := range u.score {
		if max < sc {
			max = sc
		}
	}
	return max
}
func main() {
	u := User{"Kirill", 30, []int{1, 2, 3, 97, 5, 6}}
	changeAge(&u)
	u.setName("Denis")
	fmt.Println(&u)
	fmt.Printf("Максимальный счет =%d", u.getScore())
}
