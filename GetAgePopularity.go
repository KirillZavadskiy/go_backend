package main

import "fmt"

type Person struct {
	Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	popularity := make(map[uint8]int)
	for _, p := range pl {
		popularity[p.Age]++
	}

	return popularity
}

func main() {
	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}
	fmt.Println(pl.GetAgePopularity())
}
