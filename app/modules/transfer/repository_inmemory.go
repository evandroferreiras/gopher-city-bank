package transfer

import (
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/evandroferreiras/gopher-city-bank/app/model/inmemorydb"
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
	if account == nil {
		return model.Account{}, nil
	}
	return *account, nil
}

// UpdateAccountBalance subtracts the amount of money from accountID
func (r repositoryImp) UpdateAccountBalance(id string, newBalance float64) error {
	inmemorydb.UpdateAccountBalance(id, newBalance)
	return nil
}

func (r repositoryImp) StartTransaction() error {
	return nil
}

func (r repositoryImp) CommitTransaction() {
}

func (r repositoryImp) RollbackTransaction() {
}

func (r repositoryImp) LogTransfer(transfer model.Transfer) error {
	inmemorydb.LogTransfer(transfer)
	return nil
}

func (r repositoryImp) GetAllWithdrawsOf(accountOriginID string) ([]model.Transfer, error) {
	return inmemorydb.GetAllWithdrawsOf(accountOriginID), nil
}

func (r repositoryImp) GetAllDepositsTo(accountOriginID string) ([]model.Transfer, error) {
	return inmemorydb.GetAllDepositsTo(accountOriginID), nil
}
