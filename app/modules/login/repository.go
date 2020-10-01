//go:generate mockery --name Repository --filename=repository.go

package login

import (
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/evandroferreiras/gopher-city-bank/app/model/inmemorydb"
)

// Repository is an interface to Login repository
type Repository interface {
	GetAccountByCpf(cpf string) (*model.Account, error)
}

type repositoryImp struct {
}

// NewRepository is a constructor to Login repository
func NewRepository() Repository {
	return &repositoryImp{}
}

func (r repositoryImp) GetAccountByCpf(cpf string) (*model.Account, error) {
	account := inmemorydb.GetAccountByCpf(cpf)
	return account, nil
}
