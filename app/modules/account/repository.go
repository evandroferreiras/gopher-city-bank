//go:generate mockery --name Repository --filename=repository.go

package account

import (
	"github.com/evandroferreiras/gopher-city-bank/app/common/envvar"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// Repository is an interface to Account repository
type Repository interface {
	Create(newAccount model.Account) (model.Account, error)
	GetAccounts(page int, size int) ([]model.Account, error)
	GetAccount(id string) (model.Account, error)
}

// BuildRepository is a factory constructor for Account Repository
func BuildRepository() Repository {
	if envvar.UsingMemoryDB() {
		return NewInMemoryDBRepository()
	}
	return NewORMRepository()
}
