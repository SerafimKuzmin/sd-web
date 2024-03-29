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

// CreatePerson provides a mock function with given fields: e
func (_m *RepositoryI) CreatePerson(e *models.Person) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Person) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePerson provides a mock function with given fields: id
func (_m *RepositoryI) DeletePerson(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPerson provides a mock function with given fields: id
func (_m *RepositoryI) GetPerson(id uint64) (*models.Person, error) {
	ret := _m.Called(id)

	var r0 *models.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*models.Person, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) *models.Person); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserPersons provides a mock function with given fields: userID
func (_m *RepositoryI) GetUserPersons(userID uint64) ([]*models.Person, error) {
	ret := _m.Called(userID)

	var r0 []*models.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) ([]*models.Person, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) []*models.Person); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePerson provides a mock function with given fields: e
func (_m *RepositoryI) UpdatePerson(e *models.Person) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Person) error); ok {
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
