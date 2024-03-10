package main

import (
	"fmt"
	"strings"
)

var s string = "hello world"
var mode string = "underscore"

func main() {
	fmt.Print(ModifySpaces(s, mode))
}

func ModifySpaces(s, mode string) string {
	p := ""
	switch {
	case mode == "dash":
		p = "-"
	case mode == "underscore":
		p = "_"
	default:
		p = "*"

	}
	return strings.ReplaceAll(s, " ", p)
}
