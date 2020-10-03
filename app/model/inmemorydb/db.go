package inmemorydb

import (
	"fmt"
	"sync"
	"time"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// MemoryDatabase struct
type MemoryDatabase struct {
	accounts map[int]*model.Account
}

var memoryDatabase *MemoryDatabase

var once sync.Once

// GetMemoryDB returns a map of Accounts.
func GetMemoryDB() *MemoryDatabase {
	once.Do(func() {
		memoryDatabase = &MemoryDatabase{accounts: make(map[int]*model.Account)}
	})
	return memoryDatabase
}

// AddAccount adds an account to the inMemoryDatabase
func AddAccount(newAccount model.Account) *model.Account {
	db := GetMemoryDB()
	idx := len(db.accounts)

	account := &model.Account{
		ID:        fmt.Sprintf("%d", idx+1),
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
	var accounts = make([]model.Account, 0)

	for _, account := range db.accounts {
		accounts = append(accounts, *account)
	}
	return accounts
}

// GetAccount returns an account given an id
func GetAccount(id string) *model.Account {
	db := GetMemoryDB()
	for _, account := range db.accounts {
		if account.ID == id {

			return account
		}
	}
	return nil
}

// GetAccountByCpf returns an account given an cpf
func GetAccountByCpf(cpf string) *model.Account {
	db := GetMemoryDB()
	for _, account := range db.accounts {
		if account.Cpf == cpf {
			return account
		}
	}
	return nil
}

// UpdateAccountBalance updates and account balance
func UpdateAccountBalance(id string, newBalance float64) {
	db := GetMemoryDB()
	for _, account := range db.accounts {
		if account.ID == id {
			account.Balance = newBalance
			return
		}
	}
}
