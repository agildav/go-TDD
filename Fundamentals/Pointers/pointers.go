package pointers

import "fmt"

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

// // // // // // // // // // // // // // // // // // // //

func (w *Wallet) Withdraw(amount Bitcoin) {
	if amount >= 0 {
		w.balance -= amount
	}
}

func (w *Wallet) Deposit(amount Bitcoin) {
	if amount >= 0.0 {
		w.balance += amount
	}
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// // // // // // // // // // // // // // // // // // // //

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}
