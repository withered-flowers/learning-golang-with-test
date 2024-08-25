package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

// ? Since Bitcoin is in int, we need to "parse" to string
type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

// ? When Bitcoin is parse to string, it is automatically added "n BTC"
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// ? Since we will modify the balance directly, we need pointer
// ! If not using pointer, the Wallet is using new address
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("Address of balance test is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	// ? Should we use the (*w).balance?
	// ! Nope, struct pointer is auto dereferenced in golang
	return w.balance
}
