package main

import "fmt"

// func anyPass(kyle ... interface{}){}
// if interface is passed no value accept type is reqd as it is in latter portion.
func sum(nums ... int) int{
	total:=0
	for _,num:=range nums{
		total=total+num
	}
	return total
}
func main(){
	fmt.Println(1,2,3,"Hello") 
	// in variadic func multiple values can be passed
	result:=sum(1,4,9,13,18)
	sl:=[]int{12,36,78,91,100}
	r2:=sum(sl...)
	fmt.Println(result)
	fmt.Println(r2)
}