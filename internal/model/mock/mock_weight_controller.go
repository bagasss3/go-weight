// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bagasss3/go-weight/internal/model (interfaces: WeightController)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/bagasss3/go-weight/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockWeightController is a mock of WeightController interface.
type MockWeightController struct {
	ctrl     *gomock.Controller
	recorder *MockWeightControllerMockRecorder
}

// MockWeightControllerMockRecorder is the mock recorder for MockWeightController.
type MockWeightControllerMockRecorder struct {
	mock *MockWeightController
}

// NewMockWeightController creates a new mock instance.
func NewMockWeightController(ctrl *gomock.Controller) *MockWeightController {
	mock := &MockWeightController{ctrl: ctrl}
	mock.recorder = &MockWeightControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeightController) EXPECT() *MockWeightControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWeightController) Create(arg0 context.Context, arg1 *model.WeightInput) (*model.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*model.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWeightControllerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWeightController)(nil).Create), arg0, arg1)
}

// DeleteWeight mocks base method.
func (m *MockWeightController) DeleteWeight(arg0 context.Context, arg1 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWeight", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWeight indicates an expected call of DeleteWeight.
func (mr *MockWeightControllerMockRecorder) DeleteWeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWeight", reflect.TypeOf((*MockWeightController)(nil).DeleteWeight), arg0, arg1)
}

// ReadWeights mocks base method.
func (m *MockWeightController) ReadWeights(arg0 context.Context) ([]*model.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWeights", arg0)
	ret0, _ := ret[0].([]*model.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWeights indicates an expected call of ReadWeights.
func (mr *MockWeightControllerMockRecorder) ReadWeights(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWeights", reflect.TypeOf((*MockWeightController)(nil).ReadWeights), arg0)
}

// ShowWeight mocks base method.
func (m *MockWeightController) ShowWeight(arg0 context.Context, arg1 int64) (*model.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowWeight", arg0, arg1)
	ret0, _ := ret[0].(*model.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowWeight indicates an expected call of ShowWeight.
func (mr *MockWeightControllerMockRecorder) ShowWeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowWeight", reflect.TypeOf((*MockWeightController)(nil).ShowWeight), arg0, arg1)
}

// UpdateWeight mocks base method.
func (m *MockWeightController) UpdateWeight(arg0 context.Context, arg1 int64, arg2 *model.WeightInput) (*model.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWeight", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWeight indicates an expected call of UpdateWeight.
func (mr *MockWeightControllerMockRecorder) UpdateWeight(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWeight", reflect.TypeOf((*MockWeightController)(nil).UpdateWeight), arg0, arg1, arg2)
}
