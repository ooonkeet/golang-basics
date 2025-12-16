package main

import (
	"fmt"
	"slices"
)

func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums==nil)
	fmt.Println(len(nums))
	fmt.Println(nums)
	// slice-array decl
	var nums1=make([]int,3,5) //in 2nd args length is passed and 3rd args is initial capacity
	fmt.Println(cap(nums1)) //maximum capacity of the array is represented by cap function
	// dynamic resizing can be done of capacity and length
	nums1=append(nums1,4)
	nums1=append(nums1,5)
	nums1=append(nums1,6)
	nums1=append(nums1,7)
	fmt.Println(cap(nums1))	
	// append func adds element behind
	// the capacity of slice doubles once it reaches the max limit
	fmt.Println(nums1)
	// shorthand slice declaration
	nums2:=[]int{}
	fmt.Println(nums2)
	nums2=append(nums2, 5)
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))
	var nums3=make([]int,0,5)
	nums3=append(nums3,2)
	var nums4=make([]int,len(nums3))
	fmt.Println(nums3,nums4)
	// copy func
	copy(nums4,nums3)
	fmt.Println(nums3,nums4)
	// slice func
	nums5:=[]int{1,2,3,4,5}
	fmt.Println(nums5[1:3]) //excluding last idx
	fmt.Println(nums5[:4]) //no element in left means 0th idx
	fmt.Println(nums5[2:]) //no element in right means last idx
	// slice package
	fmt.Println(slices.Equal(nums3,nums4))
	// 2d slice
	var nums6 = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums6)
}