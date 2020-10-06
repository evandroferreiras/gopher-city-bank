// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	model "github.com/evandroferreiras/gopher-city-bank/app/model"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetAllDepositsTo provides a mock function with given fields: accountOriginID, page, size
func (_m *Service) GetAllDepositsTo(accountOriginID string, page int, size int) ([]model.Transfer, error) {
	ret := _m.Called(accountOriginID, page, size)

	var r0 []model.Transfer
	if rf, ok := ret.Get(0).(func(string, int, int) []model.Transfer); ok {
		r0 = rf(accountOriginID, page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Transfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(accountOriginID, page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllWithdrawsOf provides a mock function with given fields: accountOriginID, page, size
func (_m *Service) GetAllWithdrawsOf(accountOriginID string, page int, size int) ([]model.Transfer, error) {
	ret := _m.Called(accountOriginID, page, size)

	var r0 []model.Transfer
	if rf, ok := ret.Get(0).(func(string, int, int) []model.Transfer); ok {
		r0 = rf(accountOriginID, page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Transfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(accountOriginID, page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferBetweenAccount provides a mock function with given fields: accountOriginID, accountDestinationID, amount
func (_m *Service) TransferBetweenAccount(accountOriginID string, accountDestinationID string, amount float64) (model.Account, error) {
	ret := _m.Called(accountOriginID, accountDestinationID, amount)

	var r0 model.Account
	if rf, ok := ret.Get(0).(func(string, string, float64) model.Account); ok {
		r0 = rf(accountOriginID, accountDestinationID, amount)
	} else {
		r0 = ret.Get(0).(model.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, float64) error); ok {
		r1 = rf(accountOriginID, accountDestinationID, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
