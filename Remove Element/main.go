package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
	fmt.Println(removeElement([]int{1}, 1))
}

// removeElement removes all instances of val in nums in-place and returns the new length.
func removeElement(nums []int, val int) int {
    count := 0
		for i, number := range nums {
		if number != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
		
}