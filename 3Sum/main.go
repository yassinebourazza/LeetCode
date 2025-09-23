package main

import "fmt"

func main() {
	fmt.Println(threeSum([]int{0,3,0,1,1,-1,-5,-5,3,-3,-3,0}))
	fmt.Println(threeSum([]int{0,1,1}))
	fmt.Println(threeSum([]int{0,0,0}))

}

func threeSum(nums []int) [][]int {
    var allResults [][]int
	var checker []int

	for i,n1 := range nums {
		for j,n2 := range nums {
			for k,n3 := range nums {
				if (i!=j) && (i!=k) && (j!=k) && (n1+n2+n3) == 0 {
					if repeatChecker(n1,n2,n3,allResults) {
						checker = append(checker, n1,n2,n3)
						allResults = append(allResults, checker)
						checker = []int{}
					}
				}
			}
		}
	}
	return  allResults
}

func repeatChecker(n1,n2,n3 int, allResults [][]int) bool {
	var count int
	var f1,f2,f3 bool
	for _,result := range allResults {
		count = 0
		f1 = false
		f2 = false
		f3 = false
			for _,r := range result {
				if r==n1 && !f1{
					count++
					f1 =true
				} else if r==n2 && !f2 {
					count++
					f2 = true
				} else if r==n3 && !f3 {
					count++
					f3 = true
				}
				if count == 3 {
					return false
				}
			}
		}
	
	return true
}