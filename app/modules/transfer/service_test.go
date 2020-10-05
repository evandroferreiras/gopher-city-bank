package transfer

import (
	"testing"

	serviceError "github.com/evandroferreiras/gopher-city-bank/app/common/service"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/pkg/errors"

	"github.com/evandroferreiras/gopher-city-bank/app/modules/transfer/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRepository() *mocks.Repository {
	return &mocks.Repository{}
}

func Test_TransferBetweenAccount_ShouldReturnNoErrorAndAccountOrigin_WhenSuccessfully(t *testing.T) {
	const amountToTransfer = 100

	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.
		getAccountOrigin(accountOriginReturned, nil).
		getAccountOrigin(accountOriginReturned, nil).
		getAccountDestination(accountDestinationReturned, nil).
		startTransaction(nil).
		commitTransaction(nil).
		updateAccountBalanceOrigin(400, nil).
		updateAccountBalanceDestination(100.2, nil).
		logTransfer(nil).
		build()

	service := serviceImp{repository: repositoryMock}
	accountReturned, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amountToTransfer)

	assert.NoError(t, err)
	assert.Equal(t, accountOriginID, accountReturned.ID)
	assert.Equal(t, accountOriginReturned.Balance-amountToTransfer, mockBuilder.CapturedTotalAfterWithdraw)
	assert.Equal(t, accountDestinationReturned.Balance+amountToTransfer, mockBuilder.CapturedTotalAfterDeposit)
}

func Test_TransferBetweenAccount_FloatingPoint(t *testing.T) {
	const amountToTransfer = 0.1

	accountOriginReturned := accountOriginReturned
	accountOriginReturned.Balance = 0.2
	accountDestinationReturned := accountDestinationReturned
	accountDestinationReturned.Balance = 0.2

	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.
		getAccountOrigin(accountOriginReturned, nil).
		getAccountOrigin(accountOriginReturned, nil).
		getAccountDestination(accountDestinationReturned, nil).
		startTransaction(nil).
		commitTransaction(nil).
		updateAccountBalanceOrigin(0.1, nil).
		updateAccountBalanceDestination(0.3, nil).
		logTransfer(nil).
		build()

	service := serviceImp{repository: repositoryMock}
	accountReturned, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amountToTransfer)
	assert.NoError(t, err)
	assert.Equal(t, accountOriginID, accountReturned.ID)
	t.Logf("total after withdraw: %+v total after deposit: %+v", mockBuilder.CapturedTotalAfterWithdraw, mockBuilder.CapturedTotalAfterDeposit)
	assert.Equal(t, 0.1, mockBuilder.CapturedTotalAfterWithdraw)
	assert.Equal(t, 0.3, mockBuilder.CapturedTotalAfterDeposit)

}

func Test_TransferBetweenAccount_ShouldReturnNotFoundError_WhenOriginAccountDoestExists(t *testing.T) {
	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.getAccountOrigin(emptyAccount, nil).build()

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amount)
	assert.Error(t, err)
	assert.EqualError(t, errors.Cause(err), serviceError.ErrorNotFound.Error())
	assert.Equal(t, emptyAccount, returnedAccount)
}

func Test_TransferBetweenAccount_ShouldReturnError_WhenGotErrorFromRepoWhileGettingOriginAccount(t *testing.T) {
	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.getAccountOrigin(emptyAccount, errors.New("some error")).build()

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amount)
	assert.Error(t, err)
	assert.EqualError(t, errors.Cause(err), "some error")
	assert.Equal(t, emptyAccount, returnedAccount)
}

func Test_TransferBetweenAccount_ShouldReturnNotFoundError_WhenDestinationAccountDoestExists(t *testing.T) {
	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.
		getAccountOrigin(accountOriginReturned, nil).
		getAccountDestination(emptyAccount, nil).
		build()

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amount)
	assert.Error(t, err)
	assert.EqualError(t, errors.Cause(err), serviceError.ErrorNotFound.Error())
	assert.Equal(t, emptyAccount, returnedAccount)
}

func Test_TransferBetweenAccount_ShouldReturnError_WhenGotErrorFromRepoWhileGettingDestinationAccount(t *testing.T) {
	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.
		getAccountOrigin(accountOriginReturned, nil).
		getAccountDestination(emptyAccount, errors.New("some error")).
		build()

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amount)
	assert.Error(t, err)
	assert.EqualError(t, errors.Cause(err), "some error")
	assert.Equal(t, emptyAccount, returnedAccount)
}

func Test_TransferBetweenAccount_ShouldReturnError_WhenGotErrorFromStartTransaction(t *testing.T) {
	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.
		getAccountOrigin(accountOriginReturned, nil).
		getAccountDestination(accountDestinationReturned, nil).
		startTransaction(errors.New("transaction error")).
		build()

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amount)
	assert.Error(t, err)
	assert.EqualError(t, errors.Cause(err), "transaction error")
	assert.Equal(t, emptyAccount, returnedAccount)
}

func Test_TransferBetweenAccount_ShouldReturnError_WhenThereIsNotEnoughAccountBalance(t *testing.T) {
	accountOriginReturned := accountOriginReturned
	accountOriginReturned.Balance = 0.0

	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.
		getAccountOrigin(accountOriginReturned, nil).
		getAccountDestination(accountDestinationReturned, nil).
		startTransaction(nil).
		build()

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amount)
	assert.Error(t, err)
	assert.EqualError(t, errors.Cause(err), serviceError.ErrorNotEnoughAccountBalance.Error())
	assert.Equal(t, emptyAccount, returnedAccount)
}

func Test_TransferBetweenAccount_ShouldReturnErrorAndRollback_WhenGotErrorWhileUpdatingOriginAccountBalance(t *testing.T) {
	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.
		getAccountOrigin(accountOriginReturned, nil).
		getAccountDestination(accountDestinationReturned, nil).
		startTransaction(nil).
		updateAccountBalanceOrigin(0, errors.New("update error")).
		rollbackTransaction(nil).
		build()

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amount)
	assert.EqualError(t, errors.Cause(err), "update error")
	assert.Equal(t, emptyAccount, returnedAccount)
}

func Test_TransferBetweenAccount_ShouldReturnErrorAndRollback_WhenGotErrorWhileUpdatingDestinationAccountBalance(t *testing.T) {
	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.
		getAccountOrigin(accountOriginReturned, nil).
		getAccountDestination(accountDestinationReturned, nil).
		startTransaction(nil).
		updateAccountBalanceOrigin(0, nil).
		updateAccountBalanceDestination(500.2, errors.New("update error")).
		rollbackTransaction(nil).
		build()

	service := serviceImp{repository: repositoryMock}
	returnedAccount, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amount)
	assert.EqualError(t, errors.Cause(err), "update error")
	assert.Equal(t, emptyAccount, returnedAccount)

}

func Test_TransferBetweenAccount_ShouldReturnErrorAndRollback_WhenGotErrorWhileTryingToLogTransfer(t *testing.T) {
	const amountToTransfer = 100

	mockBuilder := newMockBuilder()
	repositoryMock := mockBuilder.
		getAccountOrigin(accountOriginReturned, nil).
		getAccountOrigin(accountOriginReturned, nil).
		getAccountDestination(accountDestinationReturned, nil).
		startTransaction(nil).
		commitTransaction(nil).
		updateAccountBalanceOrigin(400, nil).
		updateAccountBalanceDestination(100.2, nil).
		logTransfer(errors.New("error on transfer")).
		build()

	service := serviceImp{repository: repositoryMock}
	accountReturned, err := service.TransferBetweenAccount(accountOriginID, accountDestinationID, amountToTransfer)

	assert.Error(t, err)
	assert.Equal(t, emptyAccount, accountReturned)
	assert.Equal(t, accountOriginID, mockBuilder.CapturedTransfer.AccountOriginID)
	assert.Equal(t, accountDestinationID, mockBuilder.CapturedTransfer.AccountDestinationID)
	assert.Equal(t, float64(amountToTransfer), mockBuilder.CapturedTransfer.Amount)
}

// Mock Builder
type mockBuilder struct {
	*mocks.Repository
	CapturedTotalAfterWithdraw float64
	CapturedTotalAfterDeposit  float64
	CapturedTransfer           model.Transfer
}

func newMockBuilder() mockBuilder {
	return mockBuilder{Repository: &mocks.Repository{}}
}

func (m *mockBuilder) getAccountOrigin(account model.Account, err error) *mockBuilder {
	m.Repository.On("GetAccount", accountOriginID).Return(account, err).Once()
	return m
}

func (m *mockBuilder) getAccountDestination(account model.Account, err error) *mockBuilder {
	m.Repository.On("GetAccount", accountDestinationID).Return(account, err).Once()
	return m
}

func (m *mockBuilder) startTransaction(err error) *mockBuilder {
	m.Repository.On("StartTransaction").Return(nil, err).Once()
	return m
}

func (m *mockBuilder) commitTransaction(err error) *mockBuilder {
	m.Repository.On("CommitTransaction", mock.Anything).Return(err).Once()
	return m
}

func (m *mockBuilder) rollbackTransaction(err error) *mockBuilder {
	m.Repository.On("RollbackTransaction", mock.Anything).Return(err).Once()
	return m
}

func (m *mockBuilder) updateAccountBalanceOrigin(newBalance float64, returnedError error) *mockBuilder {
	m.Repository.On("UpdateAccountBalance", mock.Anything, accountOriginID, newBalance).Return(returnedError).Once().
		Run(func(args mock.Arguments) {
			m.CapturedTotalAfterWithdraw = args.Get(2).(float64)
		})
	return m
}

func (m *mockBuilder) updateAccountBalanceDestination(newBalance float64, returnedError error) *mockBuilder {
	m.Repository.On("UpdateAccountBalance", mock.Anything, accountDestinationID, newBalance).Return(returnedError).Once().
		Run(func(args mock.Arguments) {
			m.CapturedTotalAfterDeposit = args.Get(2).(float64)
		})
	return m
}

func (m *mockBuilder) logTransfer(err error) *mockBuilder {
	m.Repository.On("LogTransfer", mock.Anything, mock.Anything).Return(err).Once().
		Run(func(args mock.Arguments) {
			m.CapturedTransfer = args.Get(1).(model.Transfer)
		})
	return m
}

func (m *mockBuilder) build() *mocks.Repository {
	return m.Repository
}
