// Сортировка чисел с отбором уникальных значений
package main

import (
	"fmt"
	"sort"
)

var IDs = []int64{55, 2, 5, 55, 1, 6, 5, 21}

func UniqueSortedIDs(IDs []int64) []int64 {
	sort.SliceStable(IDs, func(i, j int) bool { return IDs[i] < IDs[j] })
	uniqPointer := 0
	for i := 1; i < len(IDs); i++ {
		if IDs[uniqPointer] != IDs[i] {
			uniqPointer++
			IDs[uniqPointer] = IDs[i]
		}
	}
	return IDs[:uniqPointer+1]
}

func main() {
	fmt.Println(UniqueSortedIDs(IDs))
}
