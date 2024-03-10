package main

import (
	"fmt"
)

func IntsmaxLen(src []int, maxLen int) []int {
	sls := []int{}
	if maxLen > len(src) {
		return src
	} else if maxLen > 0 && maxLen <= len(src) {
		return src[:maxLen]
	}
	return sls
}

func main() {
	src := []int{1, 2, 5, 4}
	maxLen := 3
	fmt.Println(IntsmaxLen(src, maxLen))

}
