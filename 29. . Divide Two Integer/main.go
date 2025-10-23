package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(1, 100))
}

func divide(dividend int, divisor int) int {
	var result int
	result = dividend / divisor
	if result > math.MaxInt32 {
		result = math.MaxInt32
	}
	return result
}
