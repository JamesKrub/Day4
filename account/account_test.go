package main

import (
	"testing"
)

func TestAccount_New(t *testing.T) {
	r := New("James")
	if r.Name != "James" {
		t.Errorf("Expected James but got %s", r.Name)
	}

}
func TestAcconut_Save(t *testing.T) {
	acc := New("James")
	Save(acc)
	if Store["James"].Amount != 0 && Store["James"].Name != "James" {
		t.Errorf("Expected James and 0 but got %s and %v", Store["James"].Name, Store["James"].Amount)
	}
}

func TestFinding_Account(t *testing.T) {
	r := FindByName("James")
	if r.Name != "James" {
		t.Errorf("Expected James got %s", r.Name)
	}
}

func TestAccount_Balance(t *testing.T) {
	acc := New("Game")
	r := acc.Balance()
	if r != 0 {
		t.Errorf("Expected 0 but got %d", r)
	}
}

func TestAccount_Deposit(t *testing.T) {
	acc := New("Stang")
	acc.Deposit(1000)
	r := acc.Balance()
	if r != 1000 {
		t.Error("Expected 1000 but got ", r)
	}
}

func TestAccount_Withdraw(t *testing.T) {
	acc := New("Stang")
	acc.Deposit(1000)
	acc.Withdraw(500)
	r := acc.Balance()
	if r != 500 {
		t.Error("Expected 500 but got ", r)
	}
}
