package main

import (
	"fmt"
)

func main() {
	// s := "abcdef"
	// s1 := s[1:2]
	// fmt.Println(s1)
	a := ^(1 << 1)
	fmt.Println(a & 147)
}
