package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
