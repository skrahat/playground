package exercises

import (
	"errors"
	"log"
)

//////////////////////
// banking exercises//
//////////////////////

type IBankAccount interface {
	GetBalance() int // returns balance
	Withdraw(amount int) error
	Deposit(amount int) error
}

type TDBank struct {
	balance int
}
type RBCBank struct {
	balance int
}

func NewTDAccount() *TDBank {
	return &TDBank{
		balance: 0,
	}
}
func (b *TDBank) GetBalance() int {
	log.Printf("TD: current balance: %d", b.balance)
	return b.balance
}
func (b *TDBank) Withdraw(amount int) error {
	newBlance := b.balance - amount - 10
	if newBlance < 0 {
		return errors.New("TD: insufficient amount")
	}
	b.balance = newBlance
	return nil
}
func (b *TDBank) Deposit(deposit int) error {
	if deposit <= 0 {
		log.Fatalf("TD: value less than 1")
		return errors.New("TD: at least 1$ required to deposit")
	}
	b.balance += deposit
	log.Printf("TD: deposited successfully: %d", b.balance)
	return nil
}

// ------------------------------------------------
func NewRBCAccount() *RBCBank {
	return &RBCBank{
		balance: 0,
	}
}
func (b *RBCBank) GetBalance() int {
	log.Printf("RBC: current balance: %d", b.balance)
	return b.balance
}
func (b *RBCBank) Withdraw(amount int) error {
	newBlance := b.balance - amount - 5
	if newBlance < 0 {
		return errors.New("RBC: insufficient amount")
	}
	b.balance = newBlance
	return nil
}
func (b *RBCBank) Deposit(deposit int) error {
	if deposit <= 0 {
		log.Fatalf("RBC: value less than 1")
		return errors.New("RBC: at least 1$ required to deposit")
	}
	b.balance += deposit
	log.Printf("RBC: deposited successfully: %d", b.balance)
	return nil
}

func Exercise7() {
	var myAccounts = []IBankAccount{
		NewRBCAccount(),
		NewTDAccount(),
	}
	log.Println(myAccounts)
}
