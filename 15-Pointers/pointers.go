package main

import "fmt"

func changeNum(num int) {
	num = 5 //by value pass
	fmt.Println("In changeNum",num)
}
func changeNumByRef(num *int){
	*num=5 //by ref pass
	fmt.Println("In changeNum By ref",*num)
}
func main() {
	num:=1
	changeNum(num)
	// fmt.Println("Memory address",&num)
	fmt.Println("In main",num)
	changeNumByRef(&num)
	fmt.Println("Now in main",num)
}