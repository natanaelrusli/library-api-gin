// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// SqlGdbc is an autogenerated mock type for the SqlGdbc type
type SqlGdbc struct {
	mock.Mock
}

// ExecContext provides a mock function with given fields: ctx, query, args
func (_m *SqlGdbc) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecContext")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryContext provides a mock function with given fields: ctx, query, args
func (_m *SqlGdbc) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryContext")
	}

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sql.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryRowContext provides a mock function with given fields: ctx, query, args
func (_m *SqlGdbc) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRowContext")
	}

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Row); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// NewSqlGdbc creates a new instance of SqlGdbc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSqlGdbc(t interface {
	mock.TestingT
	Cleanup(func())
}) *SqlGdbc {
	mock := &SqlGdbc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
