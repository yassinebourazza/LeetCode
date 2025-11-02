package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1,1}))
	fmt.Println(removeDuplicates([]int{1,1,2}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}

func removeDuplicates(nums []int) int {
	slice := []int{}
	count := 0
	lenght := len(nums)
	for i:= 0 ; i < lenght; i++{
		if !contain(nums[i], slice) {
			slice = append(slice, nums[i])
			count++
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			lenght--
			i--
		}
	}
    return count
}

func contain(value int, slice []int) bool {
	for _,compare :=range slice {
		if value == compare {
			return true
		}
	}
	return  false
}