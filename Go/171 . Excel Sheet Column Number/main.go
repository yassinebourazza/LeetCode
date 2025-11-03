package main

import ("fmt"
         "math")

func main() {
	fmt.Println(titleToNumber("AB"))
}

func titleToNumber(columnTitle string) int {
    var result int
    for i,r :=range columnTitle {
          result +=  int(math.Pow(26,float64(len(columnTitle)-1-i))) * (int(r)-64)
    }
    return  result
}