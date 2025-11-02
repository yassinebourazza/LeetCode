package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var result int
  
	for i := 0 ; i < len(s); i++ {
		if s[i] == 'M' {
			result += 1000
		} else if (i != len(s)-1 && s[i:i+2] == "CM") {
			result += 900
			i++
		} else if s[i] == 'D' {
			result +=500
		} else if (i != len(s)-1 && s[i:i+2] == "CD") {
			result += 400
			i++
		} else if s[i] == 'C' {
			result += 100
		} else if (i != len(s)-1 && s[i:i+2] == "XC") {
			result += 90
			i++
		} else if s[i]== 'L' {
			result += 50 
		} else if (i != len(s)-1 && s[i:i+2] == "XL") {
			result += 40
			i++
		} else if s[i] == 'X' {
			result += 10
		} else if (i != len(s)-1 && s[i:i+2] == "IX") {
			result += 9
			i++
		} else if s[i] == 'V' {
			result += 5
		} else if (i != len(s)-1 && s[i:i+2] == "IV") {
			result += 4
			i++
		} else if s[i] == 'I' {
			result += 1
		}
	}
	return result
}
