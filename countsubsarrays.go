package main

import "fmt"

func main() {
	fmt.Println(countSubarrays([]int{4, 5, 6}))
	fmt.Println(countSubarrays([]int{0}))
	fmt.Println(countSubarrays([]int{-1, -4 , -1,4}))
}

func countSubarrays(nums []int) int {
	var count int
	if len(nums) < 3 {
		return count
	}
	for i:=2 ; i < len(nums) ;i++ {
		if float64(nums[i]+nums[i-2]) == (float64(nums[i-1])/2) {
			count++
		}
	}
	return count
}
