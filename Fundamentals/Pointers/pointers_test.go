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

	t.Run("positive deposit", func(t *testing.T) {
		wallet := Wallet{balance: 0.0}
		deposit := Bitcoin(10.0)
		want := Bitcoin(10.0)
		wallet.Deposit(deposit)
		assertBalance(t, wallet, want)
	})

	t.Run("negative deposit", func(t *testing.T) {
		wallet := Wallet{balance: 0.0}
		deposit := Bitcoin(-10.0)
		want := Bitcoin(0.0)
		wallet.Deposit(deposit)
		assertBalance(t, wallet, want)
	})
}
