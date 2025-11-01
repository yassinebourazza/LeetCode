package main

import "fmt"

func main() {
	fmt.Println(addBinary("1111","1111"))
	fmt.Println(addBinary("1010","1011"))

}

func addBinary(a string, b string) string {
    var result string
	var lenghtA = len(a)
	var lenghtB = len(b)
	if lenghtA < lenghtB {
		for i := 0 ; i< lenghtB-lenghtA; i++ {
			a = "0" + a
		}
	} else {
		for i := 0 ; i< lenghtA-lenghtB; i++ {
			b = "0" + b
		}
	}

	var flag bool
	for i := len(a)-1 ; i >= 0 ; i-- {
		if a[i] == '1' && b[i] == '1' {
			if !flag {
				result = "0" + result
				flag = true
			} else {
				result = "1" + result
			}
		} else if a[i] == '1' || b[i] == '1' {
			if flag{
				result = "0" + result
			} else {
				result = "1" + result
			}
		} else {
			if flag {
				result = "1" + result
				if i != 0 {
					flag= false
				}
			} else {
				result = "0" + result

			}
		}
	}
	if flag {
		result = "1" + result
	}

	return result
}