package main

import (
	"fmt"
)

const msg = "OK"
const (
	OK        = 0
	CANCELLED = 1
	UNKNOWN   = 2
)

func ErrorMessageToCode(msg string) int {
	result := 0
	if msg == "OK" {
		result = OK
	} else if msg == "CANCELLED" {
		result = CANCELLED
	} else {
		result = UNKNOWN
	}
	return result
}

func main() {
	fmt.Println(ErrorMessageToCode(msg))
}
