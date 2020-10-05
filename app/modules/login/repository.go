//go:generate mockery --name Repository --filename=repository.go

package login

import (
	"github.com/evandroferreiras/gopher-city-bank/app/common/envvar"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// Repository is an interface to Login repository
type Repository interface {
	GetAccountByCpf(cpf string) (*model.Account, error)
}

// BuildRepository is a factory constructor for Login Repository
func BuildRepository() Repository {
	if envvar.UsingMemoryDB() {
		return NewInMemoryDBRepository()
	}
	return NewORMRepository()
}
