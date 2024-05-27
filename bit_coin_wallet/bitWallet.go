package bitcoinwallet

import (
	"errors"
	"fmt"
)

var errorInsuficientBalance = errors.New("it's not possible withdraw: insuficient balance")

// Its possible create a new type from existing one
type BitCoin int

type Wallet struct {
	balance BitCoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(qtd BitCoin) {
	fmt.Printf("The balance address in the deposit is: %v \n", &w.balance)
	w.balance += qtd
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Withdraw(qtd BitCoin) error {
	if w.balance < qtd {
		return errorInsuficientBalance
	}

	w.balance -= qtd
	return nil
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
