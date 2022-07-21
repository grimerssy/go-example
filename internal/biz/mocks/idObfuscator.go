// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IdObfuscator is an autogenerated mock type for the IdObfuscator type
type IdObfuscator struct {
	mock.Mock
}

type IdObfuscator_Expecter struct {
	mock *mock.Mock
}

func (_m *IdObfuscator) EXPECT() *IdObfuscator_Expecter {
	return &IdObfuscator_Expecter{mock: &_m.Mock}
}

// DeobfuscateId provides a mock function with given fields: obfuscated
func (_m *IdObfuscator) DeobfuscateId(obfuscated int64) (int64, error) {
	ret := _m.Called(obfuscated)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(obfuscated)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(obfuscated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IdObfuscator_DeobfuscateId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeobfuscateId'
type IdObfuscator_DeobfuscateId_Call struct {
	*mock.Call
}

// DeobfuscateId is a helper method to define mock.On call
//  - obfuscated int64
func (_e *IdObfuscator_Expecter) DeobfuscateId(obfuscated interface{}) *IdObfuscator_DeobfuscateId_Call {
	return &IdObfuscator_DeobfuscateId_Call{Call: _e.mock.On("DeobfuscateId", obfuscated)}
}

func (_c *IdObfuscator_DeobfuscateId_Call) Run(run func(obfuscated int64)) *IdObfuscator_DeobfuscateId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *IdObfuscator_DeobfuscateId_Call) Return(_a0 int64, _a1 error) *IdObfuscator_DeobfuscateId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ObfuscateId provides a mock function with given fields: id
func (_m *IdObfuscator) ObfuscateId(id int64) (int64, error) {
	ret := _m.Called(id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IdObfuscator_ObfuscateId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ObfuscateId'
type IdObfuscator_ObfuscateId_Call struct {
	*mock.Call
}

// ObfuscateId is a helper method to define mock.On call
//  - id int64
func (_e *IdObfuscator_Expecter) ObfuscateId(id interface{}) *IdObfuscator_ObfuscateId_Call {
	return &IdObfuscator_ObfuscateId_Call{Call: _e.mock.On("ObfuscateId", id)}
}

func (_c *IdObfuscator_ObfuscateId_Call) Run(run func(id int64)) *IdObfuscator_ObfuscateId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *IdObfuscator_ObfuscateId_Call) Return(_a0 int64, _a1 error) *IdObfuscator_ObfuscateId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type NewIdObfuscatorT interface {
	mock.TestingT
	Cleanup(func())
}

// NewIdObfuscator creates a new instance of IdObfuscator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIdObfuscator(t NewIdObfuscatorT) *IdObfuscator {
	mock := &IdObfuscator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
