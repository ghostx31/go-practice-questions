package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, x int) int {
	sort.Ints(arr)
	low := arr[0]
	high := len(arr)
	for low <= high {
		mid := (high + low) / 2
		// fmt.Printf("Mid is %v and x is %v", mid, x)
		if mid == x {
			return mid
		}
		if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println("This program is meant to perform the binary search algorithm on your input.")
	var size int

	fmt.Printf("Enter the size of the array :")
	// fmt.Scanf(size)
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter the %dth element\n", i)
		fmt.Scan(&arr[i])
	}
	var x int
	fmt.Println("Enter the key you want to search: ")
	fmt.Scan(&x)
	fmt.Println(binarySearch(arr, x))

}
