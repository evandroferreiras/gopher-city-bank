package transfer

import (
	"context"

	"github.com/evandroferreiras/gopher-city-bank/app/db/inmemorydb"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

type repositoryImp struct {
}

// NewInMemoryDBRepository is a constructor to inmemoryDB Transfer repository
func NewInMemoryDBRepository() Repository {
	return &repositoryImp{}
}

// getAccount return a account given an id
func (r repositoryImp) GetAccount(id string) (model.Account, error) {
	account := inmemorydb.GetAccount(id)
	if account == emptyAccount {
		return emptyAccount, nil
	}
	return account, nil
}

// UpdateAccountBalance subtracts the amount of money from accountID
func (r repositoryImp) UpdateAccountBalance(ctx context.Context, id string, newBalance float64) error {
	inmemorydb.UpdateAccountBalance(id, newBalance)
	return nil
}

func (r repositoryImp) StartTransaction() (interface{}, error) {
	return nil, nil
}

func (r repositoryImp) CommitTransaction(ctx context.Context) {
}

func (r repositoryImp) RollbackTransaction(ctx context.Context) {
}

func (r repositoryImp) LogTransfer(ctx context.Context, transfer model.Transfer) error {
	inmemorydb.LogTransfer(transfer)
	return nil
}

func (r repositoryImp) GetAllWithdrawsOf(accountOriginID string) ([]model.Transfer, error) {
	return inmemorydb.GetAllWithdrawsOf(accountOriginID), nil
}

func (r repositoryImp) GetAllDepositsTo(accountOriginID string) ([]model.Transfer, error) {
	return inmemorydb.GetAllDepositsTo(accountOriginID), nil
}
