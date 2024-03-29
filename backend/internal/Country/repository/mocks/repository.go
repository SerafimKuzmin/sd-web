// Code generated by mockery v2.23.2. DO NOT EDIT.

package mocks

import (
	models "github.com/SerafimKuzmin/sd/backend/models"

	mock "github.com/stretchr/testify/mock"
)

// RepositoryI is an autogenerated mock type for the RepositoryI type
type RepositoryI struct {
	mock.Mock
}

// CreateCountry provides a mock function with given fields: e
func (_m *RepositoryI) CreateCountry(e *models.Country) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Country) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCountry provides a mock function with given fields: id
func (_m *RepositoryI) DeleteCountry(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCountry provides a mock function with given fields: id
func (_m *RepositoryI) GetCountry(id uint64) (*models.Country, error) {
	ret := _m.Called(id)

	var r0 *models.Country
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*models.Country, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) *models.Country); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Country)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserCountrys provides a mock function with given fields: userID
func (_m *RepositoryI) GetUserCountrys(userID uint64) ([]*models.Country, error) {
	ret := _m.Called(userID)

	var r0 []*models.Country
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) ([]*models.Country, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) []*models.Country); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Country)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCountry provides a mock function with given fields: e
func (_m *RepositoryI) UpdateCountry(e *models.Country) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Country) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepositoryI interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryI creates a new instance of RepositoryI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryI(t mockConstructorTestingTNewRepositoryI) *RepositoryI {
	mock := &RepositoryI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
