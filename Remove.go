package main

import "fmt"

var nums = []int{1, 2, 3}
var i = 2

func Remove(nums []int, i int) []int {
	if i < 0 || i > len(nums)-1 {
		return nums
	}

	nums[i] = nums[len(nums)-1]

	return nums[:len(nums)-1]
}

func main() {
	fmt.Println(Remove(nums, i))
}
