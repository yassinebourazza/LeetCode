package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3749))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	var result string
    integerToRoman := map[int]string{
		1000: "M",
		900: "CM",
		500: "D",
		400: "CD",
		100: "C",
		90: "XC",
		50: "L",
		40: "XL",
		10: "X",
		9: "IX",
		5: "V",
		4: "IV",
		1: "I",
	}
	for num > 0 {
		if num >= 1000 {
			result += integerToRoman[1000]
			num -= 1000
		} else if num >= 900 {
			result += integerToRoman[900]
			num -= 900
		}else if num >= 500 {
			result += integerToRoman[500]
			num -= 500
		}else if num >= 400 {
			result += integerToRoman[400]
			num -= 400
		}else if num >= 100 {
			result += integerToRoman[100]
			num -= 100
		}else if num >= 90 {
			result += integerToRoman[90]
			num -= 90
		}else if num >= 50 {
			result += integerToRoman[50]
			num -= 50
		}else if num >= 40 {
			result += integerToRoman[40]
			num -= 40
		}else if num >= 10 {
			result += integerToRoman[10]
			num -= 10
		}else if num >= 9 {
			result += integerToRoman[9]
			num -= 9
		}else if num >= 5 {
			result += integerToRoman[5]
			num -= 5
		}else if num >= 4 {
			result += integerToRoman[4]
			num -= 4
		}else if num >= 1 {
			result += integerToRoman[1]
			num -= 1
		}
	}


	return result
}