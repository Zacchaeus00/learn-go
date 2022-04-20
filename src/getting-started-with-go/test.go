package main

import "fmt"

func main() {
	a := []int{9, 8, 7}
	for i := range a {
		fmt.Println(i)
	}
}
