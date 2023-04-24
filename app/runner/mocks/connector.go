// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"sync"

	"github.com/umputun/simplotask/app/remote"
)

// ConnectorMock is a mock implementation of runner.Connector.
//
//	func TestSomethingThatUsesConnector(t *testing.T) {
//
//		// make and configure a mocked runner.Connector
//		mockedConnector := &ConnectorMock{
//			ConnectFunc: func(ctx context.Context, host string) (*remote.Executer, error) {
//				panic("mock out the Connect method")
//			},
//			UserFunc: func() string {
//				panic("mock out the User method")
//			},
//		}
//
//		// use mockedConnector in code that requires runner.Connector
//		// and then make assertions.
//
//	}
type ConnectorMock struct {
	// ConnectFunc mocks the Connect method.
	ConnectFunc func(ctx context.Context, host string) (*remote.Executer, error)

	// UserFunc mocks the User method.
	UserFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Connect holds details about calls to the Connect method.
		Connect []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Host is the host argument value.
			Host string
		}
		// User holds details about calls to the User method.
		User []struct {
		}
	}
	lockConnect sync.RWMutex
	lockUser    sync.RWMutex
}

// Connect calls ConnectFunc.
func (mock *ConnectorMock) Connect(ctx context.Context, host string) (*remote.Executer, error) {
	if mock.ConnectFunc == nil {
		panic("ConnectorMock.ConnectFunc: method is nil but Connector.Connect was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Host string
	}{
		Ctx:  ctx,
		Host: host,
	}
	mock.lockConnect.Lock()
	mock.calls.Connect = append(mock.calls.Connect, callInfo)
	mock.lockConnect.Unlock()
	return mock.ConnectFunc(ctx, host)
}

// ConnectCalls gets all the calls that were made to Connect.
// Check the length with:
//
//	len(mockedConnector.ConnectCalls())
func (mock *ConnectorMock) ConnectCalls() []struct {
	Ctx  context.Context
	Host string
} {
	var calls []struct {
		Ctx  context.Context
		Host string
	}
	mock.lockConnect.RLock()
	calls = mock.calls.Connect
	mock.lockConnect.RUnlock()
	return calls
}

// User calls UserFunc.
func (mock *ConnectorMock) User() string {
	if mock.UserFunc == nil {
		panic("ConnectorMock.UserFunc: method is nil but Connector.User was just called")
	}
	callInfo := struct {
	}{}
	mock.lockUser.Lock()
	mock.calls.User = append(mock.calls.User, callInfo)
	mock.lockUser.Unlock()
	return mock.UserFunc()
}

// UserCalls gets all the calls that were made to User.
// Check the length with:
//
//	len(mockedConnector.UserCalls())
func (mock *ConnectorMock) UserCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockUser.RLock()
	calls = mock.calls.User
	mock.lockUser.RUnlock()
	return calls
}