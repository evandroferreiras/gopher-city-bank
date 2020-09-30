//go:generate mockery --name Repository --filename=repository.go

package account

import (
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/evandroferreiras/gopher-city-bank/app/model/inmemorydb"
)

// Repository is an interface to Account repository
type Repository interface {
	Create(newAccount model.Account) (*model.Account, error)
	GetAccounts() ([]model.Account, error)
}

type repositoryImp struct {
}

// NewRepository is a constructor to Account repository
func NewRepository() Repository {
	return &repositoryImp{}
}

func (r *repositoryImp) Create(newAccount model.Account) (*model.Account, error) {
	accountAdded := inmemorydb.AddAccount(newAccount)

	return &accountAdded, nil
}

func (r *repositoryImp) GetAccounts() ([]model.Account, error) {
	return inmemorydb.GetAccounts(), nil
}
