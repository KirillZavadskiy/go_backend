package main

import "fmt"

func main() {
	numberLists := []int{1, 2, 3}
	fmt.Println(MergeNumberLists(numberLists))
}

func MergeNumberLists(numberLists ...[]int) []int {
	res := []int{}
	for _, n := range numberLists {
		res = append(res, n...)
	}
	return res
}
