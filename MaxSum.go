package main

import (
	"fmt"
)

func MaxSum(nums1, nums2 []int) []int {
	res1 := 0
	res2 := 0
	r := nums1
	for _, n := range nums1 {
		go func(n int) {
			res1 = res1 + n
		}(n)
	}
	for _, b := range nums2 {
		go func(b int) {
			res2 = res2 + b
		}(b)
	}
	if res1 > res2 {
		r = nums1
	} else {
		r = nums2
	}
	return r
}

func main() {
	a, b := []int{1, 2, 3}, []int{10, 20, 50}
	fmt.Println(MaxSum(a, b))

}
