package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}

func singleNumber(nums []int) int {
	var counter = make(map[int]int)
	for _, value := range nums {
		counter[value] += 1
	}

	for i :=range counter {
		if counter[i] == 1 {
			return i
		}
	}
	return 0
}
