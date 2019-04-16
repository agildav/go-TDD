package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			errMsg := fmt.Sprintf("got -> %s, want -> %s", got, want)
			t.Errorf(colorError(errMsg))
		}
	}

	// Deposit

	t.Run("positive deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0.0)}
		deposit := Bitcoin(10.0)
		want := Bitcoin(10.0)
		wallet.Deposit(deposit)
		assertBalance(t, wallet, want)
	})

	t.Run("negative deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0.0)}
		deposit := Bitcoin(-10.0)
		want := Bitcoin(0.0)
		wallet.Deposit(deposit)
		assertBalance(t, wallet, want)
	})

	// Withdraw
	t.Run("should withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10.0)}
		amount := Bitcoin(6.0)
		want := Bitcoin(4.0)
		wallet.Withdraw(amount)
		assertBalance(t, wallet, want)
	})

	t.Run("negative withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10.0)}
		amount := Bitcoin(-6.0)
		want := Bitcoin(10.0)
		wallet.Withdraw(amount)
		assertBalance(t, wallet, want)
	})
}
