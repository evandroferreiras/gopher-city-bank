package account

import (
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/evandroferreiras/gopher-city-bank/app/model/inmemorydb"
)

type repositoryImp struct {
}

// NewInMemoryDBRepository is a constructor to inmemory Account repository
func NewInMemoryDBRepository() Repository {
	return &repositoryImp{}
}

// Create a new account.
func (r *repositoryImp) Create(newAccount model.Account) (*model.Account, error) {
	accountAdded := inmemorydb.AddAccount(newAccount)
	return accountAdded, nil
}

// GetAccounts lists all accounts
func (r *repositoryImp) GetAccounts() ([]model.Account, error) {
	return inmemorydb.GetAccounts(), nil
}

// getAccount return a account given an id
func (r *repositoryImp) GetAccount(id string) (*model.Account, error) {
	account := inmemorydb.GetAccount(id)
	return account, nil
}
