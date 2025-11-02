package main

import "fmt"

func main() {
    fmt.Println(firstMissingPositive([]int{1,2,0}))
    fmt.Println(firstMissingPositive([]int{3,4,-1,1}))
    fmt.Println(firstMissingPositive([]int{7,8,9,11,12}))
}
func firstMissingPositive(nums []int) int {
    mapChecker := map[int]bool{}
    count := 0
    for _,num := range nums{
        mapChecker[num] = true
        if count < num {
            count = num
        }
    }
    i:=1
    for i <= count+1 {
        if _,exists := mapChecker[i]; !exists {
            return i
        }
        i++
    }
    return i
}