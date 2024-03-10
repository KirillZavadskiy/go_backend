package main

import (
	"fmt"
)

func UniqueIDs(userIDs []int64) []int64 {
	m, uniq := make(map[int64]struct{}), make([]int64, 0, len(userIDs))
	for _, v := range userIDs {
		if _, ok := m[v]; !ok {
			m[v], uniq = struct{}{}, append(uniq, v)
		}
	}
	return uniq
}

func main() {
	userIDs := []int64{55, 55, 33, 22, 55}
	fmt.Println(UniqueIDs(userIDs))
}
