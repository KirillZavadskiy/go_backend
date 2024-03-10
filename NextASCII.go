package main

import (
	"fmt"
)

func NextASCII(b byte) string {
	nb := b + 1
	return string(nb)
}
func PrevASCII(b byte) string {
	nb := b - 1
	return string(nb)
}

func main() {
	b := byte('s')
	fmt.Println(NextASCII(b))
	fmt.Println(PrevASCII(b))

}
