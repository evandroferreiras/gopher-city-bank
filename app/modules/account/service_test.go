package account

import (
	"testing"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/evandroferreiras/gopher-city-bank/app/modules/account/mocks"
	errors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRepository() *mocks.Repository {
	return &mocks.Repository{}
}

func Test_Create_ShouldReturnNewAccount_WhenCreateOnRepoWithSuccess(t *testing.T) {
	repositoryMock := setupRepository()
	account := &model.Account{
		ID:      1,
		Name:    "Bruce Wayne",
		Cpf:     "12345612",
		Secret:  "xxxxx",
		Balance: 1000000,
	}
	newAccount := &model.NewAccount{
		Name:    "Bruce Wayne",
		Cpf:     "12345612",
		Secret:  "xxxxx",
		Balance: 1000000,
	}
	repositoryMock.On("Create", mock.Anything).Return(account, nil)

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.Create(*newAccount)
	assert.NoError(t, err)
	assert.Equal(t, 1, returnedAccount.ID)
}

func Test_Create_ShouldReturnError_WhenCreateOnRepoWithError(t *testing.T) {
	repositoryMock := setupRepository()
	newAccount := &model.NewAccount{
		Name:    "Bruce Wayne",
		Cpf:     "12345612",
		Secret:  "xxxxx",
		Balance: 1000000,
	}
	repositoryMock.On("Create", mock.Anything).Return(nil, errors.New("Some error"))

	service := serviceImp{repository: repositoryMock}
	_, err := service.Create(*newAccount)
	assert.EqualError(t, errors.Cause(err), "Some error")
}
