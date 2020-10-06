//go:generate mockery --name Service --filename=service.go

package login

import (
	"github.com/evandroferreiras/gopher-city-bank/app/common/hash"
	"github.com/evandroferreiras/gopher-city-bank/app/common/jwt"
	"github.com/evandroferreiras/gopher-city-bank/app/common/service"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/pkg/errors"
)

const emptyToken = ""

// Service is an interface to Login service
type Service interface {
	SignIn(cpf, secret string) (string, error)
}

type serviceImp struct {
	repository Repository
}

// NewService is a constructor to Login service
func NewService() Service {
	return &serviceImp{
		repository: BuildRepository(),
	}
}

// SignIn for existing user
func (s serviceImp) SignIn(cpf, secret string) (string, error) {
	account, err := s.repository.GetAccountByCpf(cpf)
	if err != nil {
		return emptyToken, errors.New("an error occurred when trying to get account")
	}

	if account == model.EmptyAccount {
		return emptyToken, service.ErrorCpfOrSecretInvalid
	}

	receivedSecret := hash.EncryptString(secret)
	if account.Secret != receivedSecret {
		return emptyToken, service.ErrorCpfOrSecretInvalid
	}

	return generateJwtToken(account.ID)
}

func generateJwtToken(id string) (string, error) {
	return jwt.GenerateJWT(id), nil
}
