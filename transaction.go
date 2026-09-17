package main

import (
	"errors"
)

func (store *MemoryAccountStore) Deposit(id string, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	
	store.mu.Lock()
	defer store.mu.Unlock()

	acc, exists := store.accounts[id]
	if !exists {
		return errors.New("Account not found")
	}

	acc.Balance += amount
	store.accounts[id] = acc
	return nil
}

func (store *MemoryAccountStore) Withdraw(id string, amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawal amount must be greater than zero")
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	acc, exists := store.accounts[id]
	if!exists {
		return errors.New("Account not found")
	}

	if acc.Balance < amount {
		return errors.New("insuficient funds")
	}

	acc.Balance -= amount
	store.accounts[id] = acc
	return nil
}

func (store *MemoryAccountStore) Transfer(formID string, toID string, amount float64) error {
	if amount <= 0 {
		return errors.New("Transfer amount must be greater than zero")
	}

	if formID == toID {
		return errors.New("Can not transfer money to the same account")
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	fromAcc, exists := store.accounts[formID]
	if !exists {
		return errors.New("source account not found")
	}

	toAcc, exists := store.accounts[toID]
	if !exists {
		return errors.New("destination account not found")
	}

	if fromAcc.Balance < amount {
		return errors.New("insufficient funds in source account")
	}

	fromAcc.Balance -= amount
	toAcc.Balance += amount

	store.accounts[formID] = fromAcc
	store.accounts[toID] = toAcc

	return nil
}