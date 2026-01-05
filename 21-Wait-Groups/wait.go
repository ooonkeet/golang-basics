package main

import (
	"fmt"
	"sync"
)
func task(id int, w *sync.WaitGroup){ //always a pointer in receiving end
	defer w.Done() // defer keyword - similar to useEffect cleaning function
	//indicates complete of 1 goRoutine by Done keyword, increments to next goRoutine 
	fmt.Println("doing task",id)
}
func main(){
	var wg sync.WaitGroup
	for i:=0;i<=10;i++{
		wg.Add(1) //adds one go routine per iteration
		// go func(i int){
		// 	fmt.Println(i)
		// }(i)
		go task(i,&wg)
	}
	wg.Wait() //program blocks here
	fmt.Println("done")
	// inline go routine
	for j:=90;j<=100;j++{
		wg.Add(1)
		go func(j int){
			fmt.Println(j)
			defer wg.Done()
		}(j)
	}
	wg.Wait()
	fmt.Println("done")
}
// wait groups useful for proper synchronization of cpu problems