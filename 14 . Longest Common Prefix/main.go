package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
}

func longestCommonPrefix(strs []string) string {
    lens := len(strs[0])
    result := ""
    for _,word := range strs {
        if len(word) < lens {
            lens = len(word)
        }
    }

    for j:=0 ; j < lens ; j++ {
        for i,word :=range strs {
            if i == 0 {
                result += string(word[j])
            }
            if i !=0 && word[j] != result[j] {
                return result[:j]
            }
        }
    }
    return result
}