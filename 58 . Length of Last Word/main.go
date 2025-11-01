package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hell"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("a "))
}

func lengthOfLastWord(s string) int {
	if !strings.Contains(s, " ") {
		return len(s)
	}
	count := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if i != len(s)-1 && (s[i] != ' ' && s[i+1] == ' ') {
			count = i
		} else if i != len(s)-1 && (s[i] == ' ' && s[i+1] != ' ') {
			count -= i
			break
		}
		if i == 0 {
			count -= i-1
		}
	}
	return count
}
