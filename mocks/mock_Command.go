// Code generated by mockery v2.42.1. DO NOT EDIT.

package exec

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockCommand is an autogenerated mock type for the Command type
type MockCommand struct {
	mock.Mock
}

type MockCommand_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCommand) EXPECT() *MockCommand_Expecter {
	return &MockCommand_Expecter{mock: &_m.Mock}
}

// Output provides a mock function with given fields:
func (_m *MockCommand) Output() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Output")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommand_Output_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Output'
type MockCommand_Output_Call struct {
	*mock.Call
}

// Output is a helper method to define mock.On call
func (_e *MockCommand_Expecter) Output() *MockCommand_Output_Call {
	return &MockCommand_Output_Call{Call: _e.mock.On("Output")}
}

func (_c *MockCommand_Output_Call) Run(run func()) *MockCommand_Output_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_Output_Call) Return(_a0 []byte, _a1 error) *MockCommand_Output_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommand_Output_Call) RunAndReturn(run func() ([]byte, error)) *MockCommand_Output_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields:
func (_m *MockCommand) Run() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommand_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockCommand_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
func (_e *MockCommand_Expecter) Run() *MockCommand_Run_Call {
	return &MockCommand_Run_Call{Call: _e.mock.On("Run")}
}

func (_c *MockCommand_Run_Call) Run(run func()) *MockCommand_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_Run_Call) Return(_a0 error) *MockCommand_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommand_Run_Call) RunAndReturn(run func() error) *MockCommand_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockCommand) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommand_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockCommand_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockCommand_Expecter) Start() *MockCommand_Start_Call {
	return &MockCommand_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockCommand_Start_Call) Run(run func()) *MockCommand_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_Start_Call) Return(_a0 error) *MockCommand_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommand_Start_Call) RunAndReturn(run func() error) *MockCommand_Start_Call {
	_c.Call.Return(run)
	return _c
}

// StderrPipe provides a mock function with given fields:
func (_m *MockCommand) StderrPipe() (io.ReadCloser, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StderrPipe")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func() (io.ReadCloser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() io.ReadCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommand_StderrPipe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StderrPipe'
type MockCommand_StderrPipe_Call struct {
	*mock.Call
}

// StderrPipe is a helper method to define mock.On call
func (_e *MockCommand_Expecter) StderrPipe() *MockCommand_StderrPipe_Call {
	return &MockCommand_StderrPipe_Call{Call: _e.mock.On("StderrPipe")}
}

func (_c *MockCommand_StderrPipe_Call) Run(run func()) *MockCommand_StderrPipe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_StderrPipe_Call) Return(_a0 io.ReadCloser, _a1 error) *MockCommand_StderrPipe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommand_StderrPipe_Call) RunAndReturn(run func() (io.ReadCloser, error)) *MockCommand_StderrPipe_Call {
	_c.Call.Return(run)
	return _c
}

// StdinPipe provides a mock function with given fields:
func (_m *MockCommand) StdinPipe() (io.WriteCloser, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StdinPipe")
	}

	var r0 io.WriteCloser
	var r1 error
	if rf, ok := ret.Get(0).(func() (io.WriteCloser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() io.WriteCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.WriteCloser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommand_StdinPipe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StdinPipe'
type MockCommand_StdinPipe_Call struct {
	*mock.Call
}

// StdinPipe is a helper method to define mock.On call
func (_e *MockCommand_Expecter) StdinPipe() *MockCommand_StdinPipe_Call {
	return &MockCommand_StdinPipe_Call{Call: _e.mock.On("StdinPipe")}
}

func (_c *MockCommand_StdinPipe_Call) Run(run func()) *MockCommand_StdinPipe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_StdinPipe_Call) Return(_a0 io.WriteCloser, _a1 error) *MockCommand_StdinPipe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommand_StdinPipe_Call) RunAndReturn(run func() (io.WriteCloser, error)) *MockCommand_StdinPipe_Call {
	_c.Call.Return(run)
	return _c
}

// StdoutPipe provides a mock function with given fields:
func (_m *MockCommand) StdoutPipe() (io.ReadCloser, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StdoutPipe")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func() (io.ReadCloser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() io.ReadCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommand_StdoutPipe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StdoutPipe'
type MockCommand_StdoutPipe_Call struct {
	*mock.Call
}

// StdoutPipe is a helper method to define mock.On call
func (_e *MockCommand_Expecter) StdoutPipe() *MockCommand_StdoutPipe_Call {
	return &MockCommand_StdoutPipe_Call{Call: _e.mock.On("StdoutPipe")}
}

func (_c *MockCommand_StdoutPipe_Call) Run(run func()) *MockCommand_StdoutPipe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_StdoutPipe_Call) Return(_a0 io.ReadCloser, _a1 error) *MockCommand_StdoutPipe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommand_StdoutPipe_Call) RunAndReturn(run func() (io.ReadCloser, error)) *MockCommand_StdoutPipe_Call {
	_c.Call.Return(run)
	return _c
}

// Wait provides a mock function with given fields:
func (_m *MockCommand) Wait() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Wait")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommand_Wait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Wait'
type MockCommand_Wait_Call struct {
	*mock.Call
}

// Wait is a helper method to define mock.On call
func (_e *MockCommand_Expecter) Wait() *MockCommand_Wait_Call {
	return &MockCommand_Wait_Call{Call: _e.mock.On("Wait")}
}

func (_c *MockCommand_Wait_Call) Run(run func()) *MockCommand_Wait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_Wait_Call) Return(_a0 error) *MockCommand_Wait_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommand_Wait_Call) RunAndReturn(run func() error) *MockCommand_Wait_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCommand creates a new instance of MockCommand. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommand(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommand {
	mock := &MockCommand{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}