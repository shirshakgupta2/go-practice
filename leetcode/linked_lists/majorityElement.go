package main

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
    return nums[len(nums)/2]

}

func singleNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for num, count := range numMap {
		if count == 1 {
			return num
		}
	}
	return -1
}

func singleNumber2(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func singleNumber3(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		if i == len(nums)-1 || nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
