package inmemorydb

import (
	"sync"
	"time"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

var memoryDatabase map[int]model.Account
var once sync.Once

// GetAccountsMemoryDB returns a map of Accounts.
func GetAccountsMemoryDB() map[int]model.Account {
	once.Do(func() {
		memoryDatabase = make(map[int]model.Account)
	})
	return memoryDatabase
}

// AddAccountToMemoryDB adds an account to the inmemorydatabase
func AddAccountToMemoryDB(newAccount model.NewAccount) model.Account {
	db := GetAccountsMemoryDB()
	idx := len(db)

	account := model.Account{
		ID:        idx + 1,
		Name:      newAccount.Name,
		Cpf:       newAccount.Cpf,
		Secret:    newAccount.Secret,
		Balance:   newAccount.Balance,
		CreatedAt: time.Now(),
	}
	db[idx] = account
	return account
}
