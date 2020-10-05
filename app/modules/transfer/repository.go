//go:generate mockery --name Repository --filename=repository.go

package transfer

import (
	"context"

	"github.com/evandroferreiras/gopher-city-bank/app/common/envvar"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// Repository is an interface to Transfer repository
type Repository interface {
	GetAccount(id string) (model.Account, error)
	UpdateAccountBalance(ctx context.Context, id string, newBalance float64) error
	StartTransaction() (interface{}, error)
	CommitTransaction(ctx context.Context)
	RollbackTransaction(ctx context.Context)
	LogTransfer(ctx context.Context, transfer model.Transfer) error
	GetAllWithdrawsOf(accountOriginID string) ([]model.Transfer, error)
	GetAllDepositsTo(accountOriginID string) ([]model.Transfer, error)
}

// BuildRepository is a factory constructor for Transfer Repository
func BuildRepository() Repository {
	if envvar.UsingMemoryDB() {
		return NewInMemoryDBRepository()
	}
	return NewORMRepository()
}
