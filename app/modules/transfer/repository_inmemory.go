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
	if account == model.EmptyAccount {
		return model.EmptyAccount, nil
	}
	return account, nil
}

// UpdateAccountBalance subtracts the amount of money from accountID
func (r repositoryImp) UpdateAccountBalance(_ context.Context, id string, newBalance float64) error {
	inmemorydb.UpdateAccountBalance(id, newBalance)
	return nil
}

func (r repositoryImp) StartTransaction() (interface{}, error) {
	return nil, nil
}

func (r repositoryImp) CommitTransaction(context.Context) {
}

func (r repositoryImp) RollbackTransaction(context.Context) {
}

func (r repositoryImp) LogTransfer(_ context.Context, transfer model.Transfer) error {
	inmemorydb.LogTransfer(transfer)
	return nil
}

func (r repositoryImp) GetAllWithdrawsOf(accountOriginID string, page, size int) ([]model.Transfer, error) {
	return inmemorydb.GetAllWithdrawsOf(accountOriginID, page, size), nil
}

func (r repositoryImp) GetAllDepositsTo(accountOriginID string, page, size int) ([]model.Transfer, error) {
	return inmemorydb.GetAllDepositsTo(accountOriginID, page, size), nil
}
