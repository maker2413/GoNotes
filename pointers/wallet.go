package main

import "fmt"

type Wallet struct {
	balance float64
}

func (w Wallet) Deposit(amount float64) {
	fmt.Printf("address of balance in Deposit is: %p \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() float64 {
	return w.balance
}

func main() {
	wallet := Wallet{}

	wallet.Deposit(10)

	fmt.Printf("Balance: %.2f", wallet.Balance())
}
