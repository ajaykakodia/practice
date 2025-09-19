package main

import "fmt"

// strategy defined - Added a common method to implement different strategy
type PaymentStrategy interface {
	Pay(amount float64)
}

// implement concrete strategy
// Strategy A
type CreditCardPayment struct {
	CardNumber string
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using credit card [%s]\n", amount, c.CardNumber)
}

// Strategy B
type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal [%s]\n", amount, p.Email)
}

// Strategy C
type UPIPayment struct {
	UPIId string
}

func (u *UPIPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using UPI [%s]\n", amount, u.UPIId)
}

// context
type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(s PaymentStrategy) {
	p.strategy = s
}

func (p *PaymentContext) Checkout(amount float64) {
	if p.strategy == nil {
		fmt.Println("No payment strategy set!")
		return
	}
	p.strategy.Pay(amount)
}

func main() {
	payment := &PaymentContext{}

	payment.SetStrategy(&CreditCardPayment{"1234-4567-3456"})
	payment.Checkout(234.56)

	payment.SetStrategy(&UPIPayment{"8866878786@ybl"})
	payment.Checkout(10)

	payment.SetStrategy(&PayPalPayment{"ajay.yadav@email.com"})
	payment.Checkout(25.7)
}
