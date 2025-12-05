package main

import "fmt"

func main() {
	var name string = "golang"
	var kay=5
	var isAdult = false
	fmt.Println(name)
	fmt.Println(kay)
	fmt.Println(isAdult)
	// shorthand syntax:-declaration + assignment in one line
	price:=5.11
	fmt.Println(price)
	// declaration and then use 
	var isAcPresent bool
	isAcPresent = true
	fmt.Println(isAcPresent)
}

/* static typing method

		var name string = "ankit"
	used when you need to define a var before using it.
*/

// variables and constants can be declared outside scope of main but no shorthand expression outside function scope.