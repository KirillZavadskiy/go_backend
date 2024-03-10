package main

import "fmt"

func ShiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}

	shifted := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		shift := step + int(s[i])
		shifted[i] = byte(shift)
	}
	return string(shifted)
}

func main() {
	s := "abc"
	step := 1
	fmt.Println(ShiftASCII(s, step))
}
