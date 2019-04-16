package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// // // // // // // // // // // // // // // // // // // // // // // // // //

// Wallet Errors
var errInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var errNegativeWithdrawal = errors.New("cannot withdraw, negative amount")

// // // // // // // // // // // // // // // // // // // // // // // // // //

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errInsufficientFunds
	} else if amount <= 0 {
		return errNegativeWithdrawal
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Deposit(amount Bitcoin) {
	if amount >= 0.0 {
		w.balance += amount
	}
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// // // // // // // // // // // // // // // // // // // // // // // // // //

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}
