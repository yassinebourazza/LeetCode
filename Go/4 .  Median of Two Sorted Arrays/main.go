package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,3},[]int{2}))
	fmt.Println(findMedianSortedArrays([]int{1,2},[]int{3,4}))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var number float64
	number = 0

	
	slice := []float64{}

	for i := 0; i < len(nums1); i++ {
		slice = append(slice, float64(nums1[i]))
	}

	for j := 0; j < len(nums2); j++ {
		slice = append(slice, float64(nums2[j]))
	}
	
	for a := 0; a < len(slice); a++ {
		for b := 0; b < len(slice); b++ {
			if a != b && slice[a] < slice[b] {
				slice[a], slice[b] = slice[b], slice[a]
			}
		}
	}
	fmt.Println(slice)
	
	if len(slice)%2 == 0 {
		number = (slice[(len(slice)/2)-1] + slice[len(slice)/2]) / 2
	} else {
		number = slice[(len(slice)/2)]
	}

	return number
}