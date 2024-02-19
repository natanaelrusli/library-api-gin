// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// SqlGdbc is an autogenerated mock type for the SqlGdbc type
type SqlGdbc struct {
	mock.Mock
}

// Commit provides a mock function with given fields:
func (_m *SqlGdbc) Commit() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: query, args
func (_m *SqlGdbc) Exec(query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) (sql.Result, error)); ok {
		return rf(query, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) sql.Result); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Prepare provides a mock function with given fields: query
func (_m *SqlGdbc) Prepare(query string) (*sql.Stmt, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for Prepare")
	}

	var r0 *sql.Stmt
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*sql.Stmt, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(string) *sql.Stmt); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Stmt)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: query, args
func (_m *SqlGdbc) Query(query string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) (*sql.Rows, error)); ok {
		return rf(query, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *sql.Rows); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryRow provides a mock function with given fields: query, args
func (_m *SqlGdbc) QueryRow(query string, args ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRow")
	}

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *sql.Row); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *SqlGdbc) Rollback() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxEnd provides a mock function with given fields: txFunc
func (_m *SqlGdbc) TxEnd(txFunc func() error) error {
	ret := _m.Called(txFunc)

	if len(ret) == 0 {
		panic("no return value specified for TxEnd")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func() error) error); ok {
		r0 = rf(txFunc)
	} else {
		r0 = ret.Error(0)
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