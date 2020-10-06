package account

import (
	"github.com/evandroferreiras/gopher-city-bank/app/db/inmemorydb"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

type repositoryInMemory struct {
}

// NewInMemoryDBRepository is a constructor to inmemory Account repository
func NewInMemoryDBRepository() Repository {
	return &repositoryInMemory{}
}

// Create a new account.
func (r *repositoryInMemory) Create(newAccount model.Account) (model.Account, error) {
	accountAdded := inmemorydb.AddAccount(newAccount)
	return accountAdded, nil
}

// GetAccounts lists all accounts
func (r *repositoryInMemory) GetAccounts() ([]model.Account, error) {
	return inmemorydb.GetAccounts(), nil
}

// getAccount return a account given an id
func (r *repositoryInMemory) GetAccount(id string) (model.Account, error) {
	account := inmemorydb.GetAccount(id)
	return account, nil
}
