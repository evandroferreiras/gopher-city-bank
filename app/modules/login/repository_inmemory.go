package login

import (
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/evandroferreiras/gopher-city-bank/app/model/inmemorydb"
)

type repositoryImp struct {
}

// NewInMemoryDBRepository is a constructor to inmemory Login repository
func NewInMemoryDBRepository() Repository {
	return &repositoryImp{}
}

func (r repositoryImp) GetAccountByCpf(cpf string) (*model.Account, error) {
	account := inmemorydb.GetAccountByCpf(cpf)
	return account, nil
}
