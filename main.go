package main

import "fmt"

// binary search Algorithm
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			// target found
			return mid
		} else if arr[mid] < target {
			// search in the right half
			low = mid + 1
		} else {
			// search in the left half
			high = mid - 1
		}
	}

	// target not found
	return -1
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
