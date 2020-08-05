// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-gcp/pkg/internal/client (interfaces: Interface,FirewallsService,RoutesService,FirewallsListCall,RoutesListCall,FirewallsDeleteCall,RoutesDeleteCall)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	gcp "github.com/gardener/gardener-extension-provider-gcp/pkg/internal/client"
	gomock "github.com/golang/mock/gomock"
	compute "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Firewalls mocks base method.
func (m *MockInterface) Firewalls() gcp.FirewallsService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Firewalls")
	ret0, _ := ret[0].(gcp.FirewallsService)
	return ret0
}

// Firewalls indicates an expected call of Firewalls.
func (mr *MockInterfaceMockRecorder) Firewalls() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Firewalls", reflect.TypeOf((*MockInterface)(nil).Firewalls))
}

// Routes mocks base method.
func (m *MockInterface) Routes() gcp.RoutesService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Routes")
	ret0, _ := ret[0].(gcp.RoutesService)
	return ret0
}

// Routes indicates an expected call of Routes.
func (mr *MockInterfaceMockRecorder) Routes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Routes", reflect.TypeOf((*MockInterface)(nil).Routes))
}

// MockFirewallsService is a mock of FirewallsService interface.
type MockFirewallsService struct {
	ctrl     *gomock.Controller
	recorder *MockFirewallsServiceMockRecorder
}

// MockFirewallsServiceMockRecorder is the mock recorder for MockFirewallsService.
type MockFirewallsServiceMockRecorder struct {
	mock *MockFirewallsService
}

// NewMockFirewallsService creates a new mock instance.
func NewMockFirewallsService(ctrl *gomock.Controller) *MockFirewallsService {
	mock := &MockFirewallsService{ctrl: ctrl}
	mock.recorder = &MockFirewallsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirewallsService) EXPECT() *MockFirewallsServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockFirewallsService) Delete(arg0, arg1 string) gcp.FirewallsDeleteCall {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(gcp.FirewallsDeleteCall)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFirewallsServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFirewallsService)(nil).Delete), arg0, arg1)
}

// List mocks base method.
func (m *MockFirewallsService) List(arg0 string) gcp.FirewallsListCall {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(gcp.FirewallsListCall)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockFirewallsServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFirewallsService)(nil).List), arg0)
}

// MockRoutesService is a mock of RoutesService interface.
type MockRoutesService struct {
	ctrl     *gomock.Controller
	recorder *MockRoutesServiceMockRecorder
}

// MockRoutesServiceMockRecorder is the mock recorder for MockRoutesService.
type MockRoutesServiceMockRecorder struct {
	mock *MockRoutesService
}

// NewMockRoutesService creates a new mock instance.
func NewMockRoutesService(ctrl *gomock.Controller) *MockRoutesService {
	mock := &MockRoutesService{ctrl: ctrl}
	mock.recorder = &MockRoutesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoutesService) EXPECT() *MockRoutesServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRoutesService) Delete(arg0, arg1 string) gcp.RoutesDeleteCall {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(gcp.RoutesDeleteCall)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRoutesServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoutesService)(nil).Delete), arg0, arg1)
}

// List mocks base method.
func (m *MockRoutesService) List(arg0 string) gcp.RoutesListCall {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(gcp.RoutesListCall)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRoutesServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoutesService)(nil).List), arg0)
}

// MockFirewallsListCall is a mock of FirewallsListCall interface.
type MockFirewallsListCall struct {
	ctrl     *gomock.Controller
	recorder *MockFirewallsListCallMockRecorder
}

// MockFirewallsListCallMockRecorder is the mock recorder for MockFirewallsListCall.
type MockFirewallsListCallMockRecorder struct {
	mock *MockFirewallsListCall
}

// NewMockFirewallsListCall creates a new mock instance.
func NewMockFirewallsListCall(ctrl *gomock.Controller) *MockFirewallsListCall {
	mock := &MockFirewallsListCall{ctrl: ctrl}
	mock.recorder = &MockFirewallsListCallMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirewallsListCall) EXPECT() *MockFirewallsListCallMockRecorder {
	return m.recorder
}

// Pages mocks base method.
func (m *MockFirewallsListCall) Pages(arg0 context.Context, arg1 func(*compute.FirewallList) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pages indicates an expected call of Pages.
func (mr *MockFirewallsListCallMockRecorder) Pages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pages", reflect.TypeOf((*MockFirewallsListCall)(nil).Pages), arg0, arg1)
}

// MockRoutesListCall is a mock of RoutesListCall interface.
type MockRoutesListCall struct {
	ctrl     *gomock.Controller
	recorder *MockRoutesListCallMockRecorder
}

// MockRoutesListCallMockRecorder is the mock recorder for MockRoutesListCall.
type MockRoutesListCallMockRecorder struct {
	mock *MockRoutesListCall
}

// NewMockRoutesListCall creates a new mock instance.
func NewMockRoutesListCall(ctrl *gomock.Controller) *MockRoutesListCall {
	mock := &MockRoutesListCall{ctrl: ctrl}
	mock.recorder = &MockRoutesListCallMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoutesListCall) EXPECT() *MockRoutesListCallMockRecorder {
	return m.recorder
}

// Pages mocks base method.
func (m *MockRoutesListCall) Pages(arg0 context.Context, arg1 func(*compute.RouteList) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pages indicates an expected call of Pages.
func (mr *MockRoutesListCallMockRecorder) Pages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pages", reflect.TypeOf((*MockRoutesListCall)(nil).Pages), arg0, arg1)
}

// MockFirewallsDeleteCall is a mock of FirewallsDeleteCall interface.
type MockFirewallsDeleteCall struct {
	ctrl     *gomock.Controller
	recorder *MockFirewallsDeleteCallMockRecorder
}

// MockFirewallsDeleteCallMockRecorder is the mock recorder for MockFirewallsDeleteCall.
type MockFirewallsDeleteCallMockRecorder struct {
	mock *MockFirewallsDeleteCall
}

// NewMockFirewallsDeleteCall creates a new mock instance.
func NewMockFirewallsDeleteCall(ctrl *gomock.Controller) *MockFirewallsDeleteCall {
	mock := &MockFirewallsDeleteCall{ctrl: ctrl}
	mock.recorder = &MockFirewallsDeleteCallMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirewallsDeleteCall) EXPECT() *MockFirewallsDeleteCallMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockFirewallsDeleteCall) Context(arg0 context.Context) gcp.FirewallsDeleteCall {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context", arg0)
	ret0, _ := ret[0].(gcp.FirewallsDeleteCall)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockFirewallsDeleteCallMockRecorder) Context(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockFirewallsDeleteCall)(nil).Context), arg0)
}

// Do mocks base method.
func (m *MockFirewallsDeleteCall) Do(arg0 ...googleapi.CallOption) (*compute.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].(*compute.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockFirewallsDeleteCallMockRecorder) Do(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockFirewallsDeleteCall)(nil).Do), arg0...)
}

// MockRoutesDeleteCall is a mock of RoutesDeleteCall interface.
type MockRoutesDeleteCall struct {
	ctrl     *gomock.Controller
	recorder *MockRoutesDeleteCallMockRecorder
}

// MockRoutesDeleteCallMockRecorder is the mock recorder for MockRoutesDeleteCall.
type MockRoutesDeleteCallMockRecorder struct {
	mock *MockRoutesDeleteCall
}

// NewMockRoutesDeleteCall creates a new mock instance.
func NewMockRoutesDeleteCall(ctrl *gomock.Controller) *MockRoutesDeleteCall {
	mock := &MockRoutesDeleteCall{ctrl: ctrl}
	mock.recorder = &MockRoutesDeleteCallMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoutesDeleteCall) EXPECT() *MockRoutesDeleteCallMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockRoutesDeleteCall) Context(arg0 context.Context) gcp.RoutesDeleteCall {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context", arg0)
	ret0, _ := ret[0].(gcp.RoutesDeleteCall)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockRoutesDeleteCallMockRecorder) Context(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockRoutesDeleteCall)(nil).Context), arg0)
}

// Do mocks base method.
func (m *MockRoutesDeleteCall) Do(arg0 ...googleapi.CallOption) (*compute.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].(*compute.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockRoutesDeleteCallMockRecorder) Do(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockRoutesDeleteCall)(nil).Do), arg0...)
}
