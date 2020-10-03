//go:generate mockery --name Service --filename=service.go

package transfer

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/evandroferreiras/gopher-city-bank/app/common/service"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var mutex sync.Mutex

// Service is an interface to Transfer service
type Service interface {
	TransferBetweenAccount(accountOriginID string, accountDestinationID string, amount float64) (model.Account, error)
}

var emptyAccount = model.Account{}

type serviceImp struct {
	repository Repository
}

// NewService is a constructor to Transfer service
func NewService() Service {
	return &serviceImp{
		repository: NewRepository(),
	}
}

// TransferBetweenAccount sends money from origin to destination account
func (s serviceImp) TransferBetweenAccount(accountOriginID string, accountDestinationID string, amount float64) (model.Account, error) {
	mutex.Lock()
	defer mutex.Unlock()

	accountOrigin, err := s.repository.GetAccount(accountOriginID)
	if err != nil {
		return emptyAccount, errors.Wrap(err, "account origin")
	}
	if accountOrigin.ID == "" {
		return emptyAccount, errors.Wrap(service.ErrorNotFound, "account origin")
	}

	accountDestination, err := s.repository.GetAccount(accountDestinationID)
	if err != nil {
		return emptyAccount, errors.Wrap(err, "account destination")
	}
	if accountDestination.ID == "" {
		return emptyAccount, errors.Wrap(service.ErrorNotFound, "account destination")
	}

	err = s.makeTransfer(accountOrigin, accountDestination, amount)
	if err != nil {
		return emptyAccount, err
	}

	return s.repository.GetAccount(accountOriginID)

}

func (s serviceImp) makeTransfer(accountOrigin model.Account, accountDestination model.Account, amount float64) error {
	err := s.repository.StartTransaction()
	if err != nil {
		return errors.Wrap(err, "error when trying to start transfer transaction")
	}

	if amount > accountOrigin.Balance {
		return service.ErrorNotEnoughAccountBalance
	}

	logrus.Debugf("Origin:%v Destination:%v", accountOrigin.Balance, accountDestination.Balance)

	accountOrigin.Balance = truncateTwoDecimals(accountOrigin.Balance - amount)
	accountDestination.Balance = truncateTwoDecimals(accountDestination.Balance + amount)

	err = s.repository.UpdateAccountBalance(accountOrigin.ID, accountOrigin.Balance)
	if err != nil {
		s.repository.RollbackTransaction()
		return errors.Wrap(err, "error when trying to withdraw money from account origin")
	}

	err = s.repository.UpdateAccountBalance(accountDestination.ID, accountDestination.Balance)
	if err != nil {
		s.repository.RollbackTransaction()
		return errors.Wrap(err, "error when trying to deposit money to account destination")
	}

	s.repository.CommitTransaction()
	return nil
}

func truncateTwoDecimals(f float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f), 2)
	return value
}
