package main

import "fmt"

func main() {
	age := 14
	if age >= 18{
		fmt.Println("Person is an adult.")
	} else if age>=12{
		fmt.Println("Person is a teenager.")
	} else{
		fmt.Println("Person is a child.")
	}
	// using logical operator
	var role = "admin"
	var hasPermissions = true
	if role=="admin" || hasPermissions{
		fmt.Println("User has admin permissions.")
	}
	// we can declare a variable under if construct
	if age:=21;age>=18{
		fmt.Println("Person is an adult.")
	}else{
		fmt.Println("Person is not an adult.")
	}
	// go has no ternary operators
}