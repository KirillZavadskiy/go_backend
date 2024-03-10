package main

import (
	"fmt"
)

func SafeWrite(nums [5]int, i, val int) [5]int {
	if i <= len(nums)-1 {
		nums[i] = val
		fmt.Println(nums)
	} else {
		return nums
	}
	return nums
}

func main() {
	nums := [5]int{1, 2, 3}
	i := 0
	val := 9
	fmt.Println(SafeWrite(nums, i, val))
}
