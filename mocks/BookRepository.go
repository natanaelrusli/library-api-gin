// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/natanaelrusli/library-api-gin/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// CreateOne provides a mock function with given fields: ctx, book
func (_m *BookRepository) CreateOne(ctx context.Context, book domain.Book) (domain.Book, error) {
	ret := _m.Called(ctx, book)

	if len(ret) == 0 {
		panic("no return value specified for CreateOne")
	}

	var r0 domain.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Book) (domain.Book, error)); ok {
		return rf(ctx, book)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Book) domain.Book); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Get(0).(domain.Book)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Book) error); ok {
		r1 = rf(ctx, book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchAll provides a mock function with given fields: ctx
func (_m *BookRepository) FetchAll(ctx context.Context) ([]domain.Book, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FetchAll")
	}

	var r0 []domain.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Book, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Book); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchAllWithAuthor provides a mock function with given fields: ctx
func (_m *BookRepository) FetchAllWithAuthor(ctx context.Context) ([]domain.BookWithAuthor, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FetchAllWithAuthor")
	}

	var r0 []domain.BookWithAuthor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.BookWithAuthor, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.BookWithAuthor); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BookWithAuthor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *BookRepository) GetByID(ctx context.Context, id int) (domain.Book, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 domain.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (domain.Book, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) domain.Book); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Book)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBookRepository creates a new instance of BookRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookRepository {
	mock := &BookRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
