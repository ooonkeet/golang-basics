package main

// maps->associative ds
// hash, object, dict
import (
	"fmt"
	"maps"
)

func main() {
	// creating map
	m:=make(map[int]int)
	// setting an element
	m[1]=10
	m[2]=20
	m[3]=30
	// getting an element
	fmt.Println(m[1])
	// string maps
	n:=make(map[string]int)
	n["A"]=1
	n["B"]=2
	n["C"]=3
	fmt.Println(n["C"])
	// if key is not present in map then it returns zero value 0 for int "" for string and false for boolean
	fmt.Println(n["E"])
	fmt.Println(len(n))
	// deleting an element/value
	delete(n,"A")
	fmt.Println(n)
	// delete all elements
	clear(n)
	fmt.Println(n)
	// shorthand decl
	o:=map[string]int{"price":100,"quantity":20,"orders":5}
	fmt.Println(o)
	// check
	v,ok:=o["price"]
	fmt.Println(v) // for a check it is better practice to add an element before the boolean check
	if ok{
		fmt.Println("all okay")
	} else{
		fmt.Println("not okay")
	}
	fmt.Println(ok)
	// equality check
	m1:=map[int]int{1:5,2:10,3:15}
	m2:=map[int]int{1:5,2:10,3:15}
	fmt.Println(maps.Equal(m1,m2))
}