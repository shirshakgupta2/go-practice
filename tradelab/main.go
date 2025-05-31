package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4,5,3,4,5,6,7,8,9}

	result := make([]int, len(arr))

	result = bruteForce(arr)
    fmt.Println(result)

	// Using a more efficient approach
	result = efficientProduct(arr)
	fmt.Println(result)
}

func efficientProduct(arr []int) []int {
	n := len(arr)
    result := make([]int, n)


    left := make([]int, n)
    right := make([]int, n)

 
    left[0] = 1
    for i := 1; i < n; i++ {
        left[i] = left[i-1] * arr[i-1]
    }


    right[n-1] = 1
    for i := n-2; i >= 0; i-- {
        right[i] = right[i+1] * arr[i+1]
    }

    
    for i := 0; i < n; i++ {
        result[i] = left[i] * right[i]
    }
    return result

	// time complexity: O(n)
	// space complexity: O(n)
}


func bruteForce(arr []int) []int {
	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		mul := 1
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			} else {
				mul *= arr[j]	
			}
		}
		result[i] = mul
	}
	return result

	// time complexity: O(n^2)
	// space complexity: O(n)
}
