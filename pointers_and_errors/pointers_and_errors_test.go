package pointers_and_errors

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		if wallet.balance != expected {
			t.Errorf("Expected %v, but got %v", expected, wallet.balance)
		}
	}

	assertError := func(t testing.TB, actual, expected error) {
		t.Helper()
		if actual == nil {
			t.Fatalf("Expected an error but there was none")
		}

		if !errors.Is(actual, expected) {
			t.Errorf("Expected %q, but got %q", actual, expected)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		expected := Bitcoin(10)

		if err != nil {
			t.Fatal("Got an error but didn't expect one", err)
		}

		assertBalance(t, wallet, expected)
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})

}
