package main

import "fmt"

type OrderStatus int
const(
	Received OrderStatus=iota //iota = untyped int (init value = 0)
	Confirmed //value = 1
	Prepared //value=2
	Delivered //value = 3
)
type EmployeeCategory string
const(
	Developer EmployeeCategory ="developer"
	Designer EmployeeCategory = "designer"
	Tester EmployeeCategory = "tester"
	Manager EmployeeCategory = "manager"
	Director EmployeeCategory = "director"
	CEO EmployeeCategory = "CEO"
	CTO EmployeeCategory = "CTO"
	CFO EmployeeCategory = "CFO"
	COO EmployeeCategory = "COO"
	HR EmployeeCategory = "HR"
	Intern EmployeeCategory = "Intern"
)
func changeOrderSts(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}
func checkEmployee(category EmployeeCategory){
	fmt.Println("Checking employee category",category)
}
func main(){
	changeOrderSts(Received)
	changeOrderSts(Confirmed)
	changeOrderSts(Prepared)
	changeOrderSts(Delivered)
	checkEmployee(Developer)
	checkEmployee(Designer)
	checkEmployee(Tester)
	checkEmployee(Manager)
}