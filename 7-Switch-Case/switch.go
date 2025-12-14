package main

import (
	"fmt"
	"time"
)

func main() {
	// switch case
	i := 5
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Others")
	}
	// multiple cases
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //two conditions simulatneously
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	// type switch
	whoAmI:=func(i interface{}){
		switch t:=i.(type){
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		case float64:
			fmt.Println("Float")
		default:
			fmt.Println("Unknown",t)
		}
	}
	whoAmI(78)
}