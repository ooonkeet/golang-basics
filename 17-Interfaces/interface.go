package main

import "fmt"

// interface -> naming convention -> logger for login, paymenter for payment
// no explicit word like 'implements' to be used before invoking the interface materials, in Go the interfaces check the method signature with the implemented struct, therefore implicitly implemented. Even method name should be same.
type paymenter interface{
	pay(amount float32)
	refund(amount float32,account string)
}
type payment struct {
	gateway paymenter //dependency inversion -> not dependent on concrete implementation, it depends on abstraction
}

func (p payment) makePayment(amount float32) {
	razorpayPaymentGW:=razorpay{}
	stripePaymentGW:=stripe{}
	razorpayPaymentGW.pay(amount)
	stripePaymentGW.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment through razorpay",amount)
}
type stripe struct{

}
func(s stripe)pay(amount float32){
	fmt.Println("Making payment through stripe",amount)
}
type fakePayment struct{}
func(f fakePayment)pay(amount float32){
	fmt.Println("Making payment through fake gateway",amount)
}
type paypal struct{}
func(p paypal) pay(amount float32){
	fmt.Println("Making payment through paypal",amount)
}
func (p paypal) refund(amount float32,account string){

}
func main() {
	// stripePaymentGW:=stripe{}
	// fakePaymentGW:=fakePayment{}
	paypalPaymentGW:=paypal{}
	newPayment:=payment{
		// gateway: stripePaymentGW,
		// gateway: stripe{},  -> you can directly call struct without making any name convention
		gateway: paypalPaymentGW, 
	}
	newPayment.makePayment(100.00)
}