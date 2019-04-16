package pointers

import "fmt"

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	if amount >= 0.0 {
		w.balance += amount
	}
}

// // // // // // // // // //

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

// // // // // // // // // //

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}
