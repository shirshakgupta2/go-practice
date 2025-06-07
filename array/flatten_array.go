package main

func Flat2DArray(arr [][]int) []int {

	// Initialize the flattened array
	flatArr := []int{}

	// Iterate through the rows of the array
	for _, row := range arr {

		// Append the current row to the flattened array
		flatArr = append(flatArr, row...)
	}
	return flatArr
}


func flattenArray(nested []interface{}) []interface{} {
	flattened := make([]interface{}, 0)

	for _, i := range nested {
		switch i.(type) {
		case []interface{}:
			flattenedSubArray := flattenArray(i.([]interface{}))
			flattened = append(flattened, flattenedSubArray...)
		case interface{}:
			flattened = append(flattened, i)
		}
	}

	return flattened
}