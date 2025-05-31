package main

import "fmt"

func main() {
	// This is a placeholder for the main function.
	// You can implement your sorting logic here.
	// For example, you might want to sort a slice of integers or strings.
	// You can use the sort package from the standard library.
	// Example:
	// import "sort"
	// nums := []int{5, 2, 3, 1, 4}
	// sort.Ints(nums)
	// fmt.Println(nums)
	arr := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("Original array:", arr)
	bubbleSort(arr)
	fmt.Println("Sorted array (Bubble Sort):", arr)
	selectSort(arr)
	fmt.Println("Sorted array (Selection Sort):", arr)
	mergeSort(arr)
	fmt.Println("Sorted array (Merge Sort):", arr)

}

func bubbleSort(arr []int) {
	// Implement bubble sort logic here
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func selectSort(arr []int) {
	// Implement selection sort logic here
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func mergeSort(arr []int) {
	// Implement merge sort logic here
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	mergeSort(left)
	fmt.Println("Left array after sorting:", left)
	mergeSort(right)
	merge(arr, left, right)
}

func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	// Copy any remaining elements from either subarray
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
