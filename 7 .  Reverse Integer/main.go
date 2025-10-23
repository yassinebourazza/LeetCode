package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-321))
	fmt.Println(reverse(120))

}

func reverse(x int) int {
    max := 2147483647
    if x == 0 {
        return 0
    }

    sign := 1
    if x <0 {
        x = -x
        sign = -1
    }

    sum := 0
    for x > 0 {
        sum = (x%10) + sum*10
        x = x / 10
    }

    if sum > max {
        return 0
    }
    return sum*sign
}