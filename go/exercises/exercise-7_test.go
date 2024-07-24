package exercises

import (
	"testing"
)

// Helper function to create a new TD account with a balance
func createTDAccountWithBalance(balance int) *TDBank {
	account := NewTDAccount()
	account.Deposit(balance)
	return account
}

// Helper function to create a new RBC account with a balance
func createRBCAccountWithBalance(balance int) *RBCBank {
	account := NewRBCAccount()
	account.Deposit(balance)
	return account
}

// Test for TD Bank Deposit and Withdraw
func TestTDBank(t *testing.T) {
	account := createTDAccountWithBalance(100)

	// Test Deposit
	err := account.Deposit(50)
	if err != nil {
		t.Errorf("TD: Expected no error, got %v", err)
	}
	if account.GetBalance() != 150 {
		t.Errorf("TD: Expected balance 150, got %d", account.GetBalance())
	}

	// Test Withdraw
	name, err := account.Withdraw(40)
	if err != nil {
		t.Errorf("TD: Expected no error, got %v", err)
	}
	if account.GetBalance() != 100 {
		t.Errorf("TD: Expected balance 100, got %d", account.GetBalance())
	}
	if name != "TD" {
		t.Errorf("TD: Expected name 'TD', got %s", name)
	}

	// Test Withdraw with insufficient funds
	_, err = account.Withdraw(150)
	if err == nil {
		t.Errorf("TD: Expected error, got nil")
	}
}

// Test for RBC Bank Deposit and Withdraw
func TestRBCBank(t *testing.T) {
	account := createRBCAccountWithBalance(100)

	// Test Deposit
	err := account.Deposit(50)
	if err != nil {
		t.Errorf("RBC: Expected no error, got %v", err)
	}
	if account.GetBalance() != 150 {
		t.Errorf("RBC: Expected balance 150, got %d", account.GetBalance())
	}

	// Test Withdraw
	name, err := account.Withdraw(40)
	if err != nil {
		t.Errorf("RBC: Expected no error, got %v", err)
	}
	if account.GetBalance() != 105 {
		t.Errorf("RBC: Expected balance 105, got %d", account.GetBalance())
	}
	if name != "RBC" {
		t.Errorf("RBC: Expected name 'RBC', got %s", name)
	}

	// Test Withdraw with insufficient funds
	_, err = account.Withdraw(150)
	if err == nil {
		t.Errorf("RBC: Expected error, got nil")
	}
}

// Test for Exercise7 function
func TestExercise7(t *testing.T) {
	var myAccounts = []IBankAccount{
		NewRBCAccount(),
		NewTDAccount(),
	}
	for _, v := range myAccounts {
		v.Deposit(50)
		if name, err := v.Withdraw(10); err != nil {
			t.Errorf("%v: %v", name, err)
		}
		v.Deposit(20)
		if name, err := v.Withdraw(50); err != nil {
			t.Errorf("%v: %v", name, err)
		}
		balance := v.GetBalance()
		expectedBalance := 10 // Expected final balance based on the operations
		if balance != expectedBalance {
			t.Errorf("Expected balance %d, got %d", expectedBalance, balance)
		}
	}
}
