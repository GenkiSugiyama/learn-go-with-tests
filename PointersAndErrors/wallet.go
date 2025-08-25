package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

// String()を定義することでStringerインターフェースを実装できる
// String()は %s 形式の文字列がどのように出力されるかを定義する
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// 預金する的な？
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// 残高を確認する
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
