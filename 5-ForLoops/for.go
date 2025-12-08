package main

import "fmt"

// for is the only construct in go for looping
func main(){
	//while loop using for
	i:=1
	for i<= 10{
		fmt.Println(i)
		i=i+1
	}
	// infinte loop
	// for{
	// 	fmt.Println("Hello")
	// }
	// classic for loop
	for i:=0;i<5;i++{
		// break and continue can be used
		if i==2{
			continue
		}
		if i==4{
			break
		}
		fmt.Println(i)
	}
	// range in loops - variable value start from 0 and continues till 'rangeVal - 1'
	for i:=range 3{
		fmt.Println(i)
	}
}