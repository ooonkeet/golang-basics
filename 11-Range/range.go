package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// iterate via for loops
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	// iterate via range
	// with only 1 value in range it will only print the range
	for _,num:=range nums{
		fmt.Println(num)
	}
	sum:=0
	for idx,num:=range nums{
		sum+=num
		fmt.Println(idx)
	}
	fmt.Println(sum)
	// ranging a map
	m:=map[string]string{"West Bengal":"Kolkata","Maharastra":"Mumbai","Tamil Nadu":"Chennai","Telangana":"Hyderabad"}
	for key,val:=range m{
		fmt.Println(key,val)
	}
	// ranging a string
	// idx is starting byte of rune
	// in byte lang 255 is the limit
	// suppose one chr takes 300 bytes, 255 is limit for 1st char then it is shifted to 2nd char which is taken in consid for numbering it
	// char is the unicode or ascii value
	for idx,char:=range "Ankit"{
		fmt.Println(idx,char)
		fmt.Println(idx,string(char))
	}
}