package main

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount <= 0 || (w.balance <= amount) {
		return InsufficientFundsError
	} else {
		w.balance -= amount
		fmt.Printf("withdraw %s remain %s\n", amount, w.Balance())
		return nil
	}
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	fmt.Println("the balance is", w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
