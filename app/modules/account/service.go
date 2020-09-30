//go:generate mockery --name Service --filename=service.go

package account

import (
	"github.com/evandroferreiras/gopher-city-bank/app/common/hash"
	"github.com/evandroferreiras/gopher-city-bank/app/representation"
	"github.com/pkg/errors"
)

// Service is an interface to Account service
type Service interface {
	Create(representation.NewAccount) (*representation.AccountCreated, error)
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
func (s *serviceImp) Create(newAccount representation.NewAccount) (*representation.AccountCreated, error) {

	newAccount = encryptSecret(newAccount)
	account, err := s.repository.Create(newAccount.ToModel())
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred when trying to create account")
	}
	return representation.ModelToAccountCreated(account), nil
}

func encryptSecret(account representation.NewAccount) representation.NewAccount {
	valueToEncrypt := account.Secret
	encryptedString := hash.EncryptString(valueToEncrypt)
	account.Secret = encryptedString
	return account
}
