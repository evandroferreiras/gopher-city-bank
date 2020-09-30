//go:generate mockery --name Service --filename=service.go

package account

import "github.com/evandroferreiras/gopher-city-bank/app/model"

// Service is an interface to Account service
type Service interface {
	Create(*model.NewAccount) (*model.Account, error)
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
func (s *serviceImp) Create(account *model.NewAccount) (*model.Account, error) {
	return nil, nil
}
