package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()

		// nil is synonymous with null from other programming languages.
		if got != nil {
			// t.Fatal will stop all tests and print its message if called.
			t.Fatal("Didn't want an error, but got one.")
		}
	}

	assertError := func(t testing.TB, got error) {
		t.Helper()

		if got == nil {
			t.Fatal("Wanted an error, but didn't get one.")
		}
		// got.Error() gives us the string contents of the error message.
		if got.Error() != ErrInsufficientFunds.Error() {
			t.Errorf("got %s want %s", got, ErrInsufficientFunds)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})
	t.Run("Withdraw insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(40))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}
