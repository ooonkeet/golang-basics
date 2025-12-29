package main

import (
	"fmt"
	"time"
)

// order struct

type order struct{
	id string
	amount float32
	status string
	createdAt time.Time //nanosecond precision
}
type customer struct{
	name string
	contact string
}
type order1 struct{
	id string
	amount float32
	status string
	customer //struct embedding -> benefits = composition, inheritence,etc
}
// constructor
func newOrder(id string,amount float32,status string) *order{
	return &order{
		id:id,
		amount:amount,
		status:status,
	}

}
// receiver type -> (o order) pass by value
func (o order)changeSts(status string){
	o.status=status
	fmt.Println("In changeSts",o)
}
// receiver type -> (o *order) pass by reference
func (o *order)changeStsByRef(status string){
	o.status=status //auto dereferencing in struct, hence no *o.status is used.
	fmt.Println("In changeSts By ref",o)
}
// when u need to modify use ptr else u can use normal op
// o in order stands for basic naming convention, first letter in order
// get operator
func (o order) getAmt() float32{
	return o.amount
}
func main(){
	// inline struct
	language:=struct{
		name string
		isGood bool
	}{"goLang",true}
	fmt.Println(language)
	// inline struct - 2
	party:=struct{
		name string
		leader string
		isPresent bool
	}{name:"BJP",leader: "Modi",isPresent: true}
	fmt.Println(party)
	// var myOrder order=order{}
	myOrder:=order{
		id:"1",
		amount: 50.00,
		status: "received",
	}
	// partial initialization is fine as uninit value will return nil or zero value
	myOrder.createdAt=time.Now() 
	// late addition can also be done with '.' operator
	// only time.Now() will return hex value
	fmt.Println("Order struct",myOrder)
	myOrder.changeSts("paid") 
	fmt.Println("After changeSts",myOrder) //no changes as it acts like pass by value
	myOrder.changeStsByRef("confirmed")
	fmt.Println("After changeSts By ref",myOrder) //changes as it acts like pass by ref

	fmt.Println(myOrder.status) //get operator
	yourOrder:=order{
		id:"2",
		amount: 105.60,
		status:"delivered",
		createdAt: time.Now(),
	}
	myOrder.status="paid"
	fmt.Println(yourOrder)
	fmt.Println(yourOrder.getAmt())
	fmt.Println(myOrder)
	// instances are independent of each other
	theirOrder:=newOrder("3",303.75,"packed")
	fmt.Println(theirOrder)
	newCustomer:=customer{
		name:"Ankit",
		contact:"1234567890",
	}
	newOrder1:= order1{
		id:"1",
		amount: 50.00,
		status: "received",
		customer: newCustomer,
	}
	newOrder2:=order1{
		id:"2",
		amount: 105.60,
		status: "delivered",
		customer: customer{
			name: "Gujji",
			contact: "7896543210",
		},
	}
	fmt.Println(newOrder1)
	fmt.Println(newOrder2)
	newOrder1.contact="9876221003"
	newOrder2.name="Puplu"
	fmt.Println(newOrder1)
	fmt.Println(newOrder2)
	// properties can be changed
}