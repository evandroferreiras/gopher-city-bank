package account

import (
	"testing"

	serviceError "github.com/evandroferreiras/gopher-city-bank/app/common/service"

	"github.com/evandroferreiras/gopher-city-bank/app/common/hash"
	"github.com/evandroferreiras/gopher-city-bank/app/model"

	"github.com/evandroferreiras/gopher-city-bank/app/modules/account/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRepository() *mocks.Repository {
	return &mocks.Repository{}
}

func Test_Create_ShouldReturnNewAccount_WhenCreateOnRepoWithSuccess(t *testing.T) {
	repositoryMock := setupRepository()
	account := model.Account{
		ID:      "1",
		Name:    "Bruce Wayne",
		Cpf:     "12345612",
		Balance: 1000000,
	}
	newAccount := model.Account{
		Name:    "Bruce Wayne",
		Cpf:     "12345612",
		Secret:  "xxxxx",
		Balance: 1000000,
	}
	repositoryMock.On("Create", mock.Anything).Return(account, nil)

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.Create(newAccount)
	assert.NoError(t, err)
	assert.Equal(t, "1", returnedAccount.ID)
}

func Test_Create_ShouldReturnError_WhenCreateOnRepoWithError(t *testing.T) {
	repositoryMock := setupRepository()
	newAccount := model.Account{
		Name:    "Bruce Wayne",
		Cpf:     "12345612",
		Secret:  "xxxxx",
		Balance: 1000000,
	}
	repositoryMock.On("Create", mock.Anything).Return(model.EmptyAccount, errors.New("Some error"))

	service := serviceImp{repository: repositoryMock}
	_, err := service.Create(newAccount)
	assert.EqualError(t, errors.Cause(err), "Some error")
}

func Test_Create_ShouldHashSecret(t *testing.T) {
	secret := "ihatejoker"
	hashedSecret := hash.EncryptString(secret)

	repositoryMock := setupRepository()
	account := model.Account{
		ID:      "1",
		Name:    "Bruce Wayne",
		Cpf:     "12345612",
		Balance: 1000000,
	}
	newAccount := model.Account{
		Name:    "Bruce Wayne",
		Cpf:     "12345612",
		Secret:  secret,
		Balance: 1000000,
	}

	var capturedAccount model.Account
	repositoryMock.On("Create", mock.Anything).
		Run(func(args mock.Arguments) {
			capturedAccount = args.Get(0).(model.Account)
		}).
		Return(account, nil)

	service := serviceImp{repository: repositoryMock}
	_, err := service.Create(newAccount)
	assert.NoError(t, err)
	assert.Equal(t, hashedSecret, capturedAccount.Secret)
}

func Test_GetAccounts_ShouldReturnList_WhenGetFromRepoWithSuccess(t *testing.T) {
	repositoryMock := setupRepository()

	mockAccounts := []model.Account{{}}
	repositoryMock.On("GetAccounts").Return(mockAccounts, nil)

	service := serviceImp{repository: repositoryMock}
	accounts, err := service.GetAccounts()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(accounts))
}

func Test_GetAccounts_ShouldReturnError_WhenGetErrorFromRepo(t *testing.T) {
	repositoryMock := setupRepository()
	repositoryMock.On("GetAccounts").Return(nil, errors.New("some error"))

	service := serviceImp{repository: repositoryMock}
	_, err := service.GetAccounts()
	assert.EqualError(t, errors.Cause(err), "some error")
}

func Test_GetAccount_ShouldReturnAccount_WhenGetFromRepoWithSuccess(t *testing.T) {
	repositoryMock := setupRepository()
	account := model.Account{ID: "1"}
	repositoryMock.On("GetAccount", account.ID).Return(account, nil)

	service := serviceImp{repository: repositoryMock}
	retunedAccount, err := service.GetAccount(account.ID)
	assert.NoError(t, err)
	assert.Equal(t, account.ID, retunedAccount.ID)
}

func Test_GetAccount_ShouldReturnError_WhenGetErrorFromRepo(t *testing.T) {
	repositoryMock := setupRepository()
	account := model.Account{ID: "1"}
	repositoryMock.On("GetAccount", account.ID).Return(model.EmptyAccount, errors.New("some error"))

	service := serviceImp{repository: repositoryMock}
	_, err := service.GetAccount(account.ID)

	assert.EqualError(t, errors.Cause(err), "some error")
}

func Test_GetAccount_ShouldReturnNotFoundError_WhenAccountFromRepoIsNil(t *testing.T) {
	repositoryMock := setupRepository()
	repositoryMock.On("GetAccount", "1").Return(model.EmptyAccount, nil)

	service := serviceImp{repository: repositoryMock}
	_, err := service.GetAccount("1")

	assert.EqualError(t, errors.Cause(err), serviceError.ErrorNotFound.Error())

}
