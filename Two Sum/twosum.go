package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4,15}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))

}

func twoSum(nums []int, target int) []int {
    slice := []int{}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i < j && nums[i]+nums[j] == target {
				slice = append(slice, i)
				slice = append(slice, j)

			}
		}
	}

	return slice

    
}