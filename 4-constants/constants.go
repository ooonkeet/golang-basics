package main

import "fmt"

const age = 6

func main() {
	const name = "Priya"
	fmt.Println(name,"is of age",age)
	// multiple declarations
	const(
		port = 5000
		host="localhost"
	)
	fmt.Println(port,host)
}