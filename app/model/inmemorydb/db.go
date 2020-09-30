package inmemorydb

import (
	"sync"
	"time"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// MemoryDatabase struct
type MemoryDatabase struct {
	accounts map[int]model.Account
}

var memoryDatabase MemoryDatabase

var once sync.Once

// GetMemoryDB returns a map of Accounts.
func GetMemoryDB() MemoryDatabase {
	once.Do(func() {
		memoryDatabase = MemoryDatabase{accounts: make(map[int]model.Account)}
	})
	return memoryDatabase
}

// AddAccount adds an account to the inMemoryDatabase
func AddAccount(newAccount model.Account) model.Account {
	db := GetMemoryDB()
	idx := len(db.accounts)

	account := model.Account{
		ID:        idx + 1,
		Name:      newAccount.Name,
		Cpf:       newAccount.Cpf,
		Secret:    newAccount.Secret,
		Balance:   newAccount.Balance,
		CreatedAt: time.Now(),
	}
	db.accounts[idx] = account
	return account
}

// GetAccounts returns all accounts from the inMemoryDatabase
func GetAccounts() []model.Account {
	db := GetMemoryDB()

	var accounts []model.Account = make([]model.Account, 0)
	for _, account := range db.accounts {
		accounts = append(accounts, account)
	}
	return accounts
}
