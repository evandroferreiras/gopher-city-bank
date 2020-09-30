//go:generate mockery --name Repository --filename=repository.go

package account

import (
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/evandroferreiras/gopher-city-bank/app/model/inmemorydb"
)

// Repository is an interface to Account repository
type Repository interface {
	Create(newAccount model.NewAccount) (*model.AccountCreated, error)
}

type repositoryImp struct {
}

// NewRepository is a constructor to Account repository
func NewRepository() Repository {
	return &repositoryImp{}
}

func (r *repositoryImp) Create(newAccount model.NewAccount) (*model.AccountCreated, error) {

	accountAdded := inmemorydb.AddAccountToMemoryDB(newAccount)

	return &model.AccountCreated{
		ID:        accountAdded.ID,
		Name:      accountAdded.Name,
		Cpf:       accountAdded.Cpf,
		Balance:   accountAdded.Balance,
		CreatedAt: accountAdded.CreatedAt,
	}, nil
}
