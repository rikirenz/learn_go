package main

import "testing"

func TestWallet(t *testing.T) {

	assertError := func (t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted an error but I did not get one")
		}

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertBalance := func (t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func (t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
	
		assertBalance(t, wallet, want)
	
	})

	t.Run("Withdraw", func (t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func (t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})


}


