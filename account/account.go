package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type Account struct {
	Name   string
	Amount int
}

var Store = make(map[string]*Account)

// New account
func New(name string) *Account {
	acc := &Account{name, 0}
	return acc
}

// Save new accounts to storage
func Save(account *Account) {

	Store[account.Name] = account

}

func FindByName(name string) *Account {
	if val, ok := Store[name]; ok {
		return val
	}
	return nil
}

// Withdraw money
func (a *Account) Withdraw(amount int) {
	a.Amount -= amount
}

// Deposit is money added to accounts
func (a *Account) Deposit(amount int) {
	if reflect.TypeOf(amount).String() != "int" {
		log.Fatal("This is not int")
	}
	a.Amount = amount
}

// Balance of account
func (a *Account) Balance() int {
	return a.Amount
}

func main() {
	http.HandleFunc("/add", HandleNewAccount)
	http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir("./"))))
	http.ListenAndServe(":8000", nil)
}

func HandleNewAccount(w http.ResponseWriter, r *http.Request) {
	accName := r.FormValue("accName")
	acc := &Account{accName, 0}
	Save(acc)
	rs := acc.Balance()
	fmt.Fprintf(w, "<h1>Account '%s' already created, your balance is %v ฿</h1.", acc.Name, rs)
}
