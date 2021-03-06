//go:generate mockery --name Service --filename=service.go

package transfer

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/evandroferreiras/gopher-city-bank/app/common/constant"

	"github.com/evandroferreiras/gopher-city-bank/app/common/customerror"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var mutex sync.Mutex

// Service is an interface to Transfer service
type Service interface {
	TransferBetweenAccount(accountOriginID string, accountDestinationID string, amount float64) (model.Account, error)
	GetAllWithdrawsOf(accountOriginID string, page, size int) ([]model.Transfer, error)
	GetAllDepositsTo(accountOriginID string, page, size int) ([]model.Transfer, error)
}

var emptyTransfers []model.Transfer

type serviceImp struct {
	repository Repository
}

// NewService is a constructor to Transfer service
func NewService() Service {
	return &serviceImp{
		repository: BuildRepository(),
	}
}

// TransferBetweenAccount sends money from origin to destination account
func (s serviceImp) TransferBetweenAccount(accountOriginID string, accountDestinationID string, amount float64) (model.Account, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if amount < 0 {
		return model.EmptyAccount, customerror.ErrorInvalidValue
	}

	accountOrigin, err := s.getAccount(accountOriginID)
	if err != nil {
		return model.EmptyAccount, errors.Wrap(err, "account origin")
	}

	accountDestination, err := s.getAccount(accountDestinationID)
	if err != nil {
		return model.EmptyAccount, errors.Wrap(err, "account destination")
	}

	err = s.makeTransfer(accountOrigin, accountDestination, amount)
	if err != nil {
		return model.EmptyAccount, err
	}

	return s.repository.GetAccount(accountOriginID)

}

// GetAllWithdrawsOf account origin
func (s serviceImp) GetAllWithdrawsOf(accountOriginID string, page, size int) ([]model.Transfer, error) {
	_, err := s.getAccount(accountOriginID)
	if err != nil {
		return emptyTransfers, errors.Wrap(err, "account origin")
	}

	transfers, err := s.repository.GetAllWithdrawsOf(accountOriginID, page, size)
	if err != nil {
		return emptyTransfers, errors.Wrap(err, fmt.Sprintf("error when trying to get withdraws of %+v", accountOriginID))
	}

	return transfers, nil
}

// GetAllDepositsTo account origin
func (s serviceImp) GetAllDepositsTo(accountOriginID string, page, size int) ([]model.Transfer, error) {
	_, err := s.getAccount(accountOriginID)
	if err != nil {
		return emptyTransfers, errors.Wrap(err, "account origin")
	}

	transfers, err := s.repository.GetAllDepositsTo(accountOriginID, page, size)
	if err != nil {
		return emptyTransfers, errors.Wrap(err, fmt.Sprintf("error when trying to get deposits to %+v", accountOriginID))
	}

	return transfers, nil
}

// private methods
func (s serviceImp) makeTransfer(accountOrigin model.Account, accountDestination model.Account, amount float64) error {
	tx, err := s.repository.StartTransaction()
	ctx := context.WithValue(context.Background(), constant.TransactionCtxKey, tx)

	if err != nil {
		return errors.Wrap(err, "error when trying to start transfer transaction")
	}

	if amount > accountOrigin.Balance {
		return customerror.ErrorNotEnoughAccountBalance
	}

	logrus.Debugf("Origin:%v Destination:%v", accountOrigin.Balance, accountDestination.Balance)

	accountOrigin.Balance = truncateTwoDecimals(accountOrigin.Balance - amount)
	accountDestination.Balance = truncateTwoDecimals(accountDestination.Balance + amount)

	err = s.repository.UpdateAccountBalance(ctx, accountOrigin.ID, accountOrigin.Balance)
	if err != nil {
		s.repository.RollbackTransaction(ctx)
		return errors.Wrap(err, "error when trying to withdraw money from account origin")
	}

	err = s.repository.UpdateAccountBalance(ctx, accountDestination.ID, accountDestination.Balance)
	if err != nil {
		s.repository.RollbackTransaction(ctx)
		return errors.Wrap(err, "error when trying to deposit money to account destination")
	}

	transfer := model.Transfer{
		AccountOriginID:      accountOrigin.ID,
		AccountDestinationID: accountDestination.ID,
		Amount:               amount,
		CreatedAt:            time.Now(),
	}
	err = s.repository.LogTransfer(ctx, transfer)
	if err != nil {
		return errors.Wrap(err, "error when trying to register transfer log")
	}

	s.repository.CommitTransaction(ctx)
	return nil
}

func (s serviceImp) getAccount(accountID string) (model.Account, error) {
	account, err := s.repository.GetAccount(accountID)
	if err != nil {
		return model.EmptyAccount, err
	}
	if account.ID == "" {
		return model.EmptyAccount, customerror.ErrorNotFound
	}
	return account, nil
}

func truncateTwoDecimals(f float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f), 2)
	return value
}
