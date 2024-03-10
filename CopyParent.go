package main

import (
	"fmt"
)

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

var p = &Parent{
	Name: "Harry",
	Children: []Child{
		{
			Name: "Andy",
			Age:  18,
		},
	},
}

func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}
	children := make([]Child, len(p.Children))
	copy(children, p.Children)
	return Parent{
		Name:     p.Name,
		Children: children,
	}
}

func main() {
	fmt.Println(CopyParent(p))
}
