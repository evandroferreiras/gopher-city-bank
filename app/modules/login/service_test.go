package login

import (
	"testing"

	"github.com/evandroferreiras/gopher-city-bank/app/common/hash"
	"github.com/evandroferreiras/gopher-city-bank/app/common/jwt"
	serviceError "github.com/evandroferreiras/gopher-city-bank/app/common/service"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/evandroferreiras/gopher-city-bank/app/modules/login/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func setupRepository() *mocks.Repository {
	return &mocks.Repository{}
}

func Test_SignIn_ReturnJwtToken_WhenCpfAndSecretIsValid(t *testing.T) {
	cpf := "01015015055"
	secret := "secret"
	account := model.Account{ID: "980", Cpf: cpf, Secret: hash.EncryptString(secret)}
	repositoryMock := setupRepository()
	repositoryMock.On("GetAccountByCpf", cpf).Return(account, nil)

	service := serviceImp{repository: repositoryMock}
	jwtToken, err := service.SignIn(cpf, secret)
	assert.NoError(t, err)
	accountIDOnJwt, err := jwt.GetIDFromJWT(jwtToken)
	assert.NoError(t, err)
	assert.Equal(t, "980", accountIDOnJwt)
}

func Test_SignIn_ReturnError_WhenGotErrorFromRepo(t *testing.T) {
	repositoryMock := setupRepository()
	repositoryMock.On("GetAccountByCpf", "").Return(model.EmptyAccount, errors.New("some error"))

	service := serviceImp{repository: repositoryMock}
	jwtToken, err := service.SignIn("", "")
	t.Log(err)
	assert.Error(t, err)
	assert.Empty(t, jwtToken)
}

func Test_SignIn_ReturnErrorUsernameOrSecretInvalid_WhenAccountFromRepoIsNil(t *testing.T) {
	repositoryMock := setupRepository()
	repositoryMock.On("GetAccountByCpf", "").Return(model.EmptyAccount, nil)

	service := serviceImp{repository: repositoryMock}
	jwtToken, err := service.SignIn("", "")
	assert.EqualError(t, errors.Cause(err), serviceError.ErrorCpfOrSecretInvalid.Error())
	assert.Empty(t, jwtToken)
}

func Test_SignIn_ReturnErrorUsernameOrSecretInvalid_WhenSecretDoesntMatch(t *testing.T) {
	cpf := "01015015055"
	secret := "secret"
	account := model.Account{ID: "980", Cpf: cpf, Secret: hash.EncryptString(secret)}
	repositoryMock := setupRepository()
	repositoryMock.On("GetAccountByCpf", cpf).Return(account, nil)

	service := serviceImp{repository: repositoryMock}
	jwtToken, err := service.SignIn(cpf, "wrongsecret")
	assert.EqualError(t, errors.Cause(err), serviceError.ErrorCpfOrSecretInvalid.Error())
	assert.Empty(t, jwtToken)
}
