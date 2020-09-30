//go:generate mockery --name Service --filename=service.go

package account

import (
	"github.com/evandroferreiras/gopher-city-bank/app/common/hash"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/pkg/errors"
)

// Service is an interface to Account service
type Service interface {
	Create(model.Account) (*model.Account, error)
}

type serviceImp struct {
	repository Repository
}

// NewService is a constructor to Account service
func NewService() Service {
	return &serviceImp{
		repository: NewRepository(),
	}
}

// Create a new account.
func (s *serviceImp) Create(account model.Account) (*model.Account, error) {

	account = encryptSecret(account)
	createdAccount, err := s.repository.Create(account)
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred when trying to create createdAccount")
	}
	return createdAccount, nil
}

func encryptSecret(account model.Account) model.Account {
	valueToEncrypt := account.Secret
	encryptedString := hash.EncryptString(valueToEncrypt)
	account.Secret = encryptedString
	return account
}
