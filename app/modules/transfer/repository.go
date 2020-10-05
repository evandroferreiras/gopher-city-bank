//go:generate mockery --name Repository --filename=repository.go

package transfer

import (
	"github.com/evandroferreiras/gopher-city-bank/app/common/envvar"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/sirupsen/logrus"
)

// Repository is an interface to Transfer repository
type Repository interface {
	GetAccount(id string) (model.Account, error)
	UpdateAccountBalance(id string, newBalance float64) error
	StartTransaction() error
	CommitTransaction()
	RollbackTransaction()
	LogTransfer(transfer model.Transfer) error
	GetAllWithdrawsOf(accountOriginID string) ([]model.Transfer, error)
	GetAllDepositsTo(accountOriginID string) ([]model.Transfer, error)
}

// BuildRepository is a factory constructor for Transfer Repository
func BuildRepository() Repository {
	if envvar.UsingMemoryDB() {
		return NewInMemoryDBRepository()
	}
	logrus.Fatal("Repository not implemented for transfer")
	return nil
}
