package main

import "fmt"

func main() {
	// array length
	var nums [4]int
	fmt.Println(len(nums))
	//array declaration 
	nums[0]=1
	fmt.Println(nums)
	// array assumes all the index occupying the smallest value of the datatype it is declared - zeroed values, bool -> false
	var nms [3]string
	fmt.Println(nms)
	// shorthand assignment
	kahn:=[3]int{1,2,3}
	fmt.Println(kahn)
	// multidimensional array
	kiss:=[2][2]int{{1,2},{3,4}}
	fmt.Println(kiss)
	// Benefits of using array
	// fixed size, predictable
	// memory optimization
	// constant time access
}