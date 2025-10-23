package main

import "fmt"

func main() {
	
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	rotate(matrix)
	fmt.Println(matrix)
	matrix = [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	rotate(matrix)
	fmt.Println(matrix)
	matrix = [][]int{{1}}
	rotate(matrix)
	fmt.Println(matrix)
	matrix = [][]int{{1,2},{3,4}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
    	var matrixResult [][]int
	var buffer []int
	lenghtOfLine := len(matrix[0])
	for j := 0 ; j < lenghtOfLine ; j++ {
    for i := len(matrix)-1 ; i >= 0 ; i-- {
			buffer = append(buffer, matrix[i][j])
		}
		matrixResult = append(matrixResult, buffer)
		buffer = []int{}
	}
    
     for i := range matrix {
        copy(matrix[i], matrixResult[i])
    }

}