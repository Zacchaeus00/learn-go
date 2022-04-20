package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{}
	for {
		n := 0
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			return
		}
		nums = append(nums, n)
		sort.Ints(nums)
		fmt.Println(nums)
	}
}
