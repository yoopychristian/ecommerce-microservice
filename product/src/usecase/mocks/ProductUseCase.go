// Code generated by mockery v1.0.0
package mocks

import (
	usecase "ecommerce-microservice/product/src/usecase"

	mock "github.com/stretchr/testify/mock"
)

// ProductUseCase is an autogenerated mock type for the ProductUseCase type
type ProductUseCase struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *ProductUseCase) FindAll() <-chan usecase.UseCaseResult {
	ret := _m.Called()

	var r0 <-chan usecase.UseCaseResult
	if rf, ok := ret.Get(0).(func() <-chan usecase.UseCaseResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan usecase.UseCaseResult)
		}
	}

	return r0
}

// FindByCategory provides a mock function with given fields: categoryID
func (_m *ProductUseCase) FindByCategory(categoryID int) <-chan usecase.UseCaseResult {
	ret := _m.Called(categoryID)

	var r0 <-chan usecase.UseCaseResult
	if rf, ok := ret.Get(0).(func(int) <-chan usecase.UseCaseResult); ok {
		r0 = rf(categoryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan usecase.UseCaseResult)
		}
	}

	return r0
}

// FindByID provides a mock function with given fields: id
func (_m *ProductUseCase) FindByID(id int) <-chan usecase.UseCaseResult {
	ret := _m.Called(id)

	var r0 <-chan usecase.UseCaseResult
	if rf, ok := ret.Get(0).(func(int) <-chan usecase.UseCaseResult); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan usecase.UseCaseResult)
		}
	}

	return r0
}
