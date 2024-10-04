package main

import (
	"errors"
	"fmt"
)

type Bitcoin float64

// This is a Stringer interface. Stringer is defined in the fmt package and
// lets you define how your type is printed when used with the %s format string
// in prints.
func (b Bitcoin) String() string {
	// Sprintf allows us to format the String we are returning.
	return fmt.Sprintf("%.2f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// (*w).balance += amount is also valid, but is optional in Go
	// unlike other languages.
	w.balance += amount
}

// This var is declared globally so it can be used in our tests as well.
var ErrInsufficientFunds = errors.New("Error: attempting to withdraw more than balance!")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	} else {
		w.balance -= amount
		return nil
	}
}

// If we don't do w *Wallet like the deposit function then we aren't using a pointer
// to access our original object, but instead we will be given a copy of our struct.
// Since our Balance function just needs to return the balance a copy of our struct
// will work. In general it is best practice to interact with the original though.
func (w Wallet) Balance() Bitcoin {
	return w.balance
}
