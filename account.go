package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account struct {
	ID           string
	Owner 			 string
	Balance 		 float64	
}

type MemoryAccountStore struct {
	mu 		sync.RWMutex
	accounts map[string]Account
}

func NewMemoryAccountStore() *MemoryAccountStore {
	return &MemoryAccountStore{
		accounts: make(map[string]Account),
	}
}

func (store *MemoryAccountStore) AddAccount(acc Account) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.accounts[acc.ID]; exists {
		return errors.New("Account already exists")
	}

	store.accounts[acc.ID] = acc
	return nil
}

func (store *MemoryAccountStore) GetAccount(id string) (Account, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	acc, exists := store.accounts[id]
	return acc, exists
}

func CreateAccount() {
	var acc Account

	fmt.Print("Enter Your name: ")
	fmt.Scan(&acc.Owner)

	fmt.Print("Enter your Balance: ")
	fmt.Scan(&acc.Balance)
}
