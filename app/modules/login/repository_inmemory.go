package login

import (
	"github.com/evandroferreiras/gopher-city-bank/app/db/inmemorydb"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

type repositoryInMemory struct {
}

// NewInMemoryDBRepository is a constructor to inmemory Login repository
func NewInMemoryDBRepository() Repository {
	return &repositoryInMemory{}
}

func (r repositoryInMemory) GetAccountByCpf(cpf string) (*model.Account, error) {
	account := inmemorydb.GetAccountByCpf(cpf)
	return account, nil
}
