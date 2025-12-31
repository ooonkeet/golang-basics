package main

import "fmt"

func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}
func printSliceString(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}
// generic
func printSliceGeneric[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
// 2nd method of generic decl
func printSliceGeneric2[T interface{}](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
// type specific
func printSliceGeneric3[T int | float32](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}
// comparable
func printSliceGeneric4[T comparable](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}
// multiple pass
func printSliceGeneric5[T comparable,V string](items[]T,name V){
	for _, item := range items {
		fmt.Println(item,name)
	}
}
// lifo
type stack[T any] struct{
	elements[]T
}

func main() {
	nums:=[]int{1,2,3,4,5}
	name:=[]string{"Ankit","Deeksha","Kiran","Sukriti"}
	options:=[]bool{true,false}
	values:=[]float32{4.56,9.87}
	// printSlice([]int {1,2,3})
	printSlice(nums)
	printSliceString(name)
	printSliceGeneric(options)
	printSliceGeneric2(values)
	printSliceGeneric3(nums)
	printSliceGeneric3(values)
	printSliceGeneric4(options)
	printSliceGeneric5(nums,"Cappucino")
	myStack:=stack[int]{
		elements: []int{1,2,3,4,5},
	}
	yourStack:=stack[string]{
		elements: []string{"Ankit","Deeksha","Kiran","Hemanta"},
	}
	fmt.Println(myStack)
	fmt.Println(yourStack)
}