package main

import "fmt"
// if we dont write int at end, it would default it void
func add(a int, b int) int {
	return a + b
}
// 2nd type of decl
func sub(a,b int) int{
	return a-b
}
// multiple return values
func getLang()(string,string,string,string,string,string,string){
	return "Java","C++","TypeScript","Python","Javascript","Ruby","PHP"
}
// passing fn within a fn
func combine(fn func(a,b int)int){
	result:=fn(5,3)
	fmt.Println(result)
}
// returning fn from fn
func Process() func(a int)int{
	return func(a int)int{
		return a*a
	}
}
func main() {
	result := add(5, 3)
	res2:=sub(5,3)
	fmt.Println(result,res2)
	type1,type2,type3,_,_,_,_:=getLang()
	_,_,_,type4,type5,type6,type7:=getLang()
	fmt.Println("Statically Typed = ",type1,type2,type3)
	fmt.Println("Dynamically Typed = ",type4,type5,type6,type7)
	fmt.Println(getLang())
	combine(add)
	ft:=Process()
	fmt.Println(ft(5))
}