package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrorInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != expected {
		t.Errorf("expected %s but got %s", expected, got)
	}
}
