// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/IBookService.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	models "aeon-grpc/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBookService is a mock of BookService interface.
type MockBookService struct {
	ctrl     *gomock.Controller
	recorder *MockBookServiceMockRecorder
}

// MockBookServiceMockRecorder is the mock recorder for MockBookService.
type MockBookServiceMockRecorder struct {
	mock *MockBookService
}

// NewMockBookService creates a new mock instance.
func NewMockBookService(ctrl *gomock.Controller) *MockBookService {
	mock := &MockBookService{ctrl: ctrl}
	mock.recorder = &MockBookServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookService) EXPECT() *MockBookServiceMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookService) CreateBook(arg0 models.Book) (models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0)
	ret0, _ := ret[0].(models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookServiceMockRecorder) CreateBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookService)(nil).CreateBook), arg0)
}

// DeleteBook mocks base method.
func (m *MockBookService) DeleteBook(arg0 models.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookServiceMockRecorder) DeleteBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookService)(nil).DeleteBook), arg0)
}

// GetBook mocks base method.
func (m *MockBookService) GetBook(arg0 string) (models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBook", arg0)
	ret0, _ := ret[0].(models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockBookServiceMockRecorder) GetBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockBookService)(nil).GetBook), arg0)
}

// UpdateBook mocks base method.
func (m *MockBookService) UpdateBook(arg0 models.Book) (models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", arg0)
	ret0, _ := ret[0].(models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookServiceMockRecorder) UpdateBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookService)(nil).UpdateBook), arg0)
}
