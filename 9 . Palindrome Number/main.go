package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
        if x >= 0 && x < 10 {
		return true
	}
	
	if x < 0 {
		return false
	}

	var slice []int

	for x > 0 {
		number := x % 10 
		x /= 10
		slice = append(slice, number)
	}
	for i := 0 ; i < len(slice)/2 ; i++ {
		if slice[i] != slice[len(slice)-1-i] {
			return false
		}
	}
	return true
}