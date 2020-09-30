package healthcheck

import (
	"errors"
	"testing"

	"github.com/evandroferreiras/gopher-city-bank/app/modules/healthcheck/mocks"
	"github.com/stretchr/testify/assert"
)

func setupRepository() *mocks.Repository {
	return &mocks.Repository{}
}

func Test_Service_IsWorking_ShouldReturnTrue_IfRepositoryIsPinging(t *testing.T) {
	repositoryMock := setupRepository()
	repositoryMock.On("Ping").Return(true, nil)

	service := &serviceImp{
		repository: repositoryMock,
	}

	result := service.IsWorking()

	assert.Equal(t, true, result)
}

func Test_Service_IsWorking_ShouldReturnFalse_IfRepositoryIsNotPinging(t *testing.T) {
	repositoryMock := setupRepository()
	repositoryMock.On("Ping").Return(false, nil)

	service := &serviceImp{
		repository: repositoryMock,
	}

	result := service.IsWorking()

	assert.Equal(t, false, result)
}

func Test_Service_IsWorking_ShouldReturnFalse_IfRepositoryPingReturnAnError(t *testing.T) {
	repositoryMock := setupRepository()
	repositoryMock.On("Ping").Return(true, errors.New("Some error"))

	service := &serviceImp{
		repository: repositoryMock,
	}

	result := service.IsWorking()

	assert.Equal(t, false, result)
}
