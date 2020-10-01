//go:generate mockery --name Repository --filename=repository.go

package account

import (
	"fmt"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/evandroferreiras/gopher-city-bank/app/model/inmemorydb"
	"github.com/pkg/errors"
)

// Repository is an interface to Account repository
type Repository interface {
	Create(newAccount model.Account) (*model.Account, error)
	GetAccounts() ([]model.Account, error)
	GetAccount(id string) (*model.Account, error)
}

type repositoryImp struct {
}

// NewRepository is a constructor to Account repository
func NewRepository() Repository {
	return &repositoryImp{}
}

// Create a new account.
func (r *repositoryImp) Create(newAccount model.Account) (*model.Account, error) {
	accountAdded := inmemorydb.AddAccount(newAccount)
	return &accountAdded, nil
}

// GetAccounts lists all accounts
func (r *repositoryImp) GetAccounts() ([]model.Account, error) {
	return inmemorydb.GetAccounts(), nil
}

// GetAccount return a account given an id
func (r *repositoryImp) GetAccount(id string) (*model.Account, error) {
	account := inmemorydb.GetAccount(id)
	if account == nil {
		return nil, errors.New(fmt.Sprintf("no account found to id %v", id))
	}
	return account, nil
}
