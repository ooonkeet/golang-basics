package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("doing tasks",id)
}
func main(){
	for i:=0;i<=10;i++{
		// task(i)
		go task(i) //lightweight thread 'go' keyword-> runs concurrently
		// golang passes it to go scheduler - go scheduler
		// non - blocking schedule
	}
	// inline go routine
	for j:=90;j<=100;j++{
		go func(){
			fmt.Println("do tasks",j)
		}() //example of closure
		go func(j int){
			fmt.Println(j)
		}(j) //example of receive data
	}
	time.Sleep(time.Second*2) // stop the time so that the os doesnot skip any iteration
	// concurrently runs the program - thus it follows no loop and prints in happasat order. 
	// go routines used to control cpu work
}