package pointers

import (
	"testing"
)

func TestWal(t *testing.T) {
	t.Run("Wallet - Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		// got := wallet.Balance()

		// ? Since this is pointer we will check the address
		// fmt.Printf("Address of balance test is %p \n", &wallet.balance)
		// want := Bitcoin(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Wallet - Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		// got := wallet.Balance()
		// want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Wallet - Withdraw - Insufficient Fund", func(t *testing.T) {
		// startingBalance := Bitcoin(20)
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))

	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("Wanted no error but get one ...")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Wanted error but didn't get one ...")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
