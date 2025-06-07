package main
import "fmt"
func main() {

   // Initialize the multi-dimensional array
	   arr := []interface{}{
		[]interface{}{1, 2, 3},
		[]interface{}{4, 5, 6},
		7, 8, 9,
		10, 11, 12,
		[]interface{}{13, 14, 15},
	}
	flatArr := flattenArray(arr)
	fmt.Println("The flattened array obtained from above multidimensional array is:", flatArr)
	
	// Flatten the array using the append function
	
	array := Flat2DArray([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
		{13, 14, 15},
	})
	fmt.Println("The flattened array obtained from above 2-D array is:", array)
}
