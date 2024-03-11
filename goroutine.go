// Наглядное применение горутин
package main

import (
	"fmt"
	"time"
)

func SumGo(nums1, nums2 []int) []int {
	s1Ch := make(chan int)
	go sumParallel1(nums1, s1Ch)

	s2Ch := make(chan int)
	go sumParallel1(nums2, s2Ch)

	s1 := <-s1Ch
	s2 := <-s2Ch

	if s1 > s2 {
		return nums1
	}
	return nums2
}

func Sum(nums1, nums2 []int) []int {
	s1 := sumParallel2(nums1)
	s2 := sumParallel2(nums2)
	if s1 > s2 {
		return nums1
	}
	return nums2
}

func sumParallel1(nums []int, resCh chan int) {
	s := 0
	time.Sleep(time.Second)
	for _, n := range nums {
		s += n
	}
	resCh <- s
}
func sumParallel2(nums []int) int {
	s := 0
	time.Sleep(time.Second)
	for _, n := range nums {
		s += n
	}
	return s
}

func main() {
	var time1 = time.Now()
	fmt.Printf("Сумма чисел больше в слайсе %d\n", SumGo([]int{1, 20, 3}, []int{2, 3, 3}))
	fmt.Printf("Время выполнения с горутиной %f сек.\n\n", time.Since(time1).Seconds())

	var time2 = time.Now()
	fmt.Printf("Сумма чисел больше в слайсе %d\n", Sum([]int{1, 20, 3}, []int{2, 30, 3}))
	fmt.Printf("Время выполнения без горутины %f сек.\n", time.Since(time2).Seconds())

}
