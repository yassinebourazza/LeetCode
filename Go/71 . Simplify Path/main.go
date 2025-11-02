package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/.../a/../b/c/../d/./"))


}

func simplifyPath(path string) string {
	var result string
	var stack []string
    slice := strings.Split(path, "/")	
	for i:= 0 ; i < len(slice) ; i++ {
		if slice[i] == "" || slice[i] == "." {
			continue
		}
		if slice[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			} else {
				stack = append(stack, slice[i])
		}
	}

	if len(stack) == 0 {
		return "/"
	}


	for _,dir := range stack {
		result += "/" + dir
	}

	return  result
	
} 