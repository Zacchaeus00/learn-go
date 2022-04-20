package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	fmt.Scanln(&s)
	s = strings.ToLower(s)
	if s[:1] == "i" && s[len(s)-1:] == "n" && strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
