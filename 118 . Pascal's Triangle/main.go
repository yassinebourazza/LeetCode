package main

import "fmt"

func main() {
	fmt.Println(generate(3))
	fmt.Println(generate(5))
	fmt.Println(generate(9))
}

func generate(numRows int) [][]int {
	var result [][]int
    for i := 0 ; i < numRows; i++ {
		if i == 0 {
			result = append(result, []int{1})
		} else if i == 1 {
			result = append(result, []int{1,1})
		} else {
			var arr []int
			arr= append(arr, 1)
			for j:= 1 ; j <i ; j++ {
				arr = append(arr, result[i-1][j]+result[i-1][j-1])
			}
			arr = append(arr, 1)
			result = append(result, arr)
		}
	}
	return result
}