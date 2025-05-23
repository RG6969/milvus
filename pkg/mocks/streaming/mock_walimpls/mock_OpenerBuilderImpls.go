// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock_walimpls

import (
	walimpls "github.com/milvus-io/milvus/pkg/v2/streaming/walimpls"
	mock "github.com/stretchr/testify/mock"
)

// MockOpenerBuilderImpls is an autogenerated mock type for the OpenerBuilderImpls type
type MockOpenerBuilderImpls struct {
	mock.Mock
}

type MockOpenerBuilderImpls_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOpenerBuilderImpls) EXPECT() *MockOpenerBuilderImpls_Expecter {
	return &MockOpenerBuilderImpls_Expecter{mock: &_m.Mock}
}

// Build provides a mock function with no fields
func (_m *MockOpenerBuilderImpls) Build() (walimpls.OpenerImpls, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 walimpls.OpenerImpls
	var r1 error
	if rf, ok := ret.Get(0).(func() (walimpls.OpenerImpls, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() walimpls.OpenerImpls); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(walimpls.OpenerImpls)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOpenerBuilderImpls_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type MockOpenerBuilderImpls_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
func (_e *MockOpenerBuilderImpls_Expecter) Build() *MockOpenerBuilderImpls_Build_Call {
	return &MockOpenerBuilderImpls_Build_Call{Call: _e.mock.On("Build")}
}

func (_c *MockOpenerBuilderImpls_Build_Call) Run(run func()) *MockOpenerBuilderImpls_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockOpenerBuilderImpls_Build_Call) Return(_a0 walimpls.OpenerImpls, _a1 error) *MockOpenerBuilderImpls_Build_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOpenerBuilderImpls_Build_Call) RunAndReturn(run func() (walimpls.OpenerImpls, error)) *MockOpenerBuilderImpls_Build_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with no fields
func (_m *MockOpenerBuilderImpls) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockOpenerBuilderImpls_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockOpenerBuilderImpls_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockOpenerBuilderImpls_Expecter) Name() *MockOpenerBuilderImpls_Name_Call {
	return &MockOpenerBuilderImpls_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockOpenerBuilderImpls_Name_Call) Run(run func()) *MockOpenerBuilderImpls_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockOpenerBuilderImpls_Name_Call) Return(_a0 string) *MockOpenerBuilderImpls_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOpenerBuilderImpls_Name_Call) RunAndReturn(run func() string) *MockOpenerBuilderImpls_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOpenerBuilderImpls creates a new instance of MockOpenerBuilderImpls. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOpenerBuilderImpls(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOpenerBuilderImpls {
	mock := &MockOpenerBuilderImpls{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
