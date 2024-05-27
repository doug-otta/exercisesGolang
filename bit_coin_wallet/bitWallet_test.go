package bitcoinwallet

import (
	"testing"
)

func TestBitWallet(t *testing.T) {
	errorConfirm := func(t *testing.T, result error, expected error) {
		t.Helper()
		if result == nil {
			t.Fatal("I was expecting an error, but not happened")
		}

		if result != expected {
			t.Errorf("RESULT %s - EXPECTED %s", result, expected)
		}
	}

	balanceConfirm := func(t *testing.T, wallet Wallet, expected BitCoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expected {
			t.Errorf("RESULT: %s, EXPECTED: %s", result, expected)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))

		balanceConfirm(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		err := wallet.Withdraw(BitCoin(100))

		balanceConfirm(t, wallet, BitCoin(10))
		errorConfirm(t, err, errorInsuficientBalance)
	})

	t.Run("Overdraft", func(t *testing.T) {
		inicialBalance := BitCoin(20)
		wallet := Wallet{inicialBalance}

		err := wallet.Withdraw(BitCoin(100))

		balanceConfirm(t, wallet, inicialBalance)
		errorConfirm(t, err, errorInsuficientBalance)
	})
}
