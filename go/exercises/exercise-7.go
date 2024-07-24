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
	Withdraw(amount int) (string, error)
	Deposit(amount int) error
}

type TDBank struct {
	name    string
	balance int
}
type RBCBank struct {
	name    string
	balance int
}

func NewTDAccount() *TDBank {
	return &TDBank{
		name:    "TD",
		balance: 0,
	}
}
func (b *TDBank) GetBalance() int {
	log.Printf("TD: current balance: %d", b.balance)
	return b.balance
}
func (b *TDBank) Withdraw(amount int) (string, error) {
	newBlance := b.balance - amount - 10
	if newBlance < 0 {
		return b.name, errors.New("insufficient amount")
	}
	b.balance = newBlance
	log.Printf("TD: withdrawn successfully: %d", b.balance)
	return b.name, nil
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
		name:    "RBC",
		balance: 0,
	}
}
func (b *RBCBank) GetBalance() int {
	log.Printf("RBC: current balance: %d", b.balance)
	return b.balance
}
func (b *RBCBank) Withdraw(amount int) (string, error) {
	newBlance := b.balance - amount - 5
	if newBlance < 0 {
		return b.name, errors.New("insufficient amount")
	}
	b.balance = newBlance
	log.Printf("RBC: withdrawn successfully: %d", b.balance)
	return b.name, nil
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
	for _, v := range myAccounts {
		v.Deposit(50)
		if name, err := v.Withdraw(10); err != nil {
			log.Printf("%v : %s", name, err)
		}
		v.Deposit(20)
		if name, err := v.Withdraw(50); err != nil {
			log.Printf("%v : %s", name, err)
		}
		log.Println(v.GetBalance())
	}

}
