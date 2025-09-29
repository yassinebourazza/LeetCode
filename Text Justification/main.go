package main

import "fmt"

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	fmt.Println(fullJustify([]string{"What","must","be","acknowledgment","shall","be"}, 16))
	fmt.Println(fullJustify([]string{"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"}, 20))

	fmt.Println(fullJustify([]string{"Listen","to","many,","speak","to","a","few."}, 6))
	fmt.Println(fullJustify([]string{"a"}, 1))
	fmt.Println(fullJustify([]string{"a","b","c","d","e"}, 3))
}

func fullJustify(words []string, maxWidth int) []string {
 return []string{}    
}