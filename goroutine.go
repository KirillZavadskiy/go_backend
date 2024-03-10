package main

import (
	"fmt"
)

func Sum(nums1, nums2 []int) int {
	s1Ch := make(chan int)
	go sumParallel(nums1, s1Ch)

	s2Ch := make(chan int)
	go sumParallel(nums2, s2Ch)

	s1 := <-s1Ch
	s2 := <-s2Ch

	if s1 > s2 {
		return s1
	}
	return s2
}

func sumParallel(nums []int, resCh chan int) {
	s := 0
	for _, n := range nums {
		s += n
	}
	resCh <- s
}

func main() {
	fmt.Println(Sum([]int{1, 20, 3}, []int{2, 3, 3}))
}
