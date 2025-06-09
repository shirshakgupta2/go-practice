package main

import (
	"fmt"
)

func main() {

	// Example usage of search algorithms

	//linearSearch
	// Linear search is suitable for unsorted arrays
	arr := []int{10, 20, 30, 40, 50}
	target := 30
	linearSearch(arr, target)
	//binarySearch
	// Binary search requires sorted arrays
	// It is more efficient than linear search for large sorted arrays
	// It has a time complexity of O(log n)
	// It works by repeatedly dividing the search interval in half
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target = 5
	binarySearch(arr, target)
	// array if not srted but rotated // arr = []int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
	// target = 3
	// binarySearch(arr, target) // This will not work as binary search requires sorted arrays
	// Note: For rotated sorted arrays, a modified binary search is required
	// However, the above example is a simple binary search on a sorted array
	// For rotated sorted arrays, you would typically find the pivot point and then perform binary search on the appropriate half
	// If the array is rotated, you would need to implement a modified binary search

	rotatedArr := []int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
	rotatedTarget := 3
	// For rotated sorted arrays, you would typically find the pivot point and then perform binary search on the appropriate half
	// For example, you can implement a function to find the pivot and then use binary search on the appropriate half
	binarySearchOnRotatedArray(rotatedArr, rotatedTarget)

	//jumpSearch
	// Jump search is suitable for sorted arrays
	// It works by jumping ahead by a fixed number of steps and then performing a linear search
	// It has a time complexity of O(√n)
	// It is more efficient than linear search for large sorted arrays
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target = 7
	jumpSearch(arr, target)
	//exponentialSearch
	// Exponential search is suitable for sorted arrays
	// It works by finding the range where the target may exist and then performing a binary search
	// It has a time complexity of O(log n)
	// It is more efficient than binary search for large sorted arrays
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target = 8
	exponentialSearch(arr, target)

}

func linearSearch(arr []int, target int) {
	for i, v := range arr {
		if v == target {
			fmt.Println("Linear Search: Found", target, "at index", i)
			return
		}
	}
	fmt.Println("Linear Search: Target", target, "not found")
}

func binarySearch(arr []int, target int) {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			fmt.Println("Binary Search: Found", target, "at index", mid)
			return
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Binary Search: Target", target, "not found")
}

func binarySearchOnRotatedArray(arr []int, target int) {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			fmt.Println("Binary Search on Rotated Array: Found", target, "at index", mid)
			return
		}

		// Check if the left half is sorted
		if arr[left] <= arr[mid] {
			// If target is in the left half
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// If target is in the right half
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	fmt.Println("Binary Search on Rotated Array: Target", target, "not found")
}

func jumpSearch(arr []int, target int) {
	n := len(arr)
	step := int(float64(n) / 2.0) // Jump step size
	prev := 0

	for arr[min(step, n)-1] < target {
		prev = step
		step += int(float64(n) / 2.0)
		if prev >= n {
			fmt.Println("Jump Search: Target", target, "not found")
			return
		}
	}

	for i := prev; i < min(step, n); i++ {
		if arr[i] == target {
			fmt.Println("Jump Search: Found", target, "at index", i)
			return
		}
	}
	fmt.Println("Jump Search: Target", target, "not found")
}

func exponentialSearch(arr []int, target int) {
	if arr[0] == target {
		fmt.Println("Exponential Search: Found", target, "at index 0")
		return
	}

	index := 1
	for index < len(arr) && arr[index] <= target {
		index *= 2
	}

	left := index / 2
	right := min(index, len(arr)-1)

	for i := left; i <= right; i++ {
		if arr[i] == target {
			fmt.Println("Exponential Search: Found", target, "at index", i)
			return
		}
	}
	fmt.Println("Exponential Search: Target", target, "not found")
}
