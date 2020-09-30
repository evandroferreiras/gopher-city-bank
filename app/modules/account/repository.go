//go:generate mockery --name Repository --filename=repository.go

package account

import (
	"time"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// Repository is an interface to Account repository
type Repository interface {
	Create(newAccount model.NewAccount) (*model.Account, error)
}

type repositoryImp struct {
}

// NewRepository is a constructor to Account repository
func NewRepository() Repository {
	return &repositoryImp{}
}

func (r *repositoryImp) Create(newAccount model.NewAccount) (*model.Account, error) {
	db := model.GetAccountsMemoryDB()
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

	return &account, nil
}
