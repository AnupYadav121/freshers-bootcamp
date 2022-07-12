// Code generated by MockGen. DO NOT EDIT.
// Source: 7thJulyQuestion2/Utils (interfaces: InterfaceDB)

// Package Utils is a generated GoMock package.
package Utils

import (
	Models "7thJulyQuestion2/Models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterfaceDB is a mock of InterfaceDB interface.
type MockInterfaceDB struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceDBMockRecorder
}

// MockInterfaceDBMockRecorder is the mock recorder for MockInterfaceDB.
type MockInterfaceDBMockRecorder struct {
	mock *MockInterfaceDB
}

// NewMockInterfaceDB creates a new mock instance.
func NewMockInterfaceDB(ctrl *gomock.Controller) *MockInterfaceDB {
	mock := &MockInterfaceDB{ctrl: ctrl}
	mock.recorder = &MockInterfaceDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceDB) EXPECT() *MockInterfaceDBMockRecorder {
	return m.recorder
}

// DoCreate mocks base method.
func (m *MockInterfaceDB) DoCreate(arg0 *Models.Student) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DoCreate", arg0)
}

// DoCreate indicates an expected call of DoCreate.
func (mr *MockInterfaceDBMockRecorder) DoCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoCreate", reflect.TypeOf((*MockInterfaceDB)(nil).DoCreate), arg0)
}

// DoCreateMark mocks base method.
func (m *MockInterfaceDB) DoCreateMark(arg0 *Models.SubjectMarks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DoCreateMark", arg0)
}

// DoCreateMark indicates an expected call of DoCreateMark.
func (mr *MockInterfaceDBMockRecorder) DoCreateMark(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoCreateMark", reflect.TypeOf((*MockInterfaceDB)(nil).DoCreateMark), arg0)
}

// DoDelete mocks base method.
func (m *MockInterfaceDB) DoDelete(arg0 *Models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoDelete indicates an expected call of DoDelete.
func (mr *MockInterfaceDBMockRecorder) DoDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDelete", reflect.TypeOf((*MockInterfaceDB)(nil).DoDelete), arg0)
}

// DoDeleteMark mocks base method.
func (m *MockInterfaceDB) DoDeleteMark(arg0 *Models.SubjectMarks) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDeleteMark", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoDeleteMark indicates an expected call of DoDeleteMark.
func (mr *MockInterfaceDBMockRecorder) DoDeleteMark(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDeleteMark", reflect.TypeOf((*MockInterfaceDB)(nil).DoDeleteMark), arg0)
}

// DoFind mocks base method.
func (m *MockInterfaceDB) DoFind(arg0 *[]Models.Student) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DoFind", arg0)
}

// DoFind indicates an expected call of DoFind.
func (mr *MockInterfaceDBMockRecorder) DoFind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoFind", reflect.TypeOf((*MockInterfaceDB)(nil).DoFind), arg0)
}

// DoFindMarks mocks base method.
func (m *MockInterfaceDB) DoFindMarks(arg0 *[]Models.SubjectMarks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DoFindMarks", arg0)
}

// DoFindMarks indicates an expected call of DoFindMarks.
func (mr *MockInterfaceDBMockRecorder) DoFindMarks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoFindMarks", reflect.TypeOf((*MockInterfaceDB)(nil).DoFindMarks), arg0)
}

// DoUpdate mocks base method.
func (m *MockInterfaceDB) DoUpdate(arg0 *Models.Student, arg1 Models.UpdatedStudent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DoUpdate", arg0, arg1)
}

// DoUpdate indicates an expected call of DoUpdate.
func (mr *MockInterfaceDBMockRecorder) DoUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoUpdate", reflect.TypeOf((*MockInterfaceDB)(nil).DoUpdate), arg0, arg1)
}

// DoUpdateMark mocks base method.
func (m *MockInterfaceDB) DoUpdateMark(arg0 *Models.SubjectMarks, arg1 Models.UpdatedSubjectMarks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DoUpdateMark", arg0, arg1)
}

// DoUpdateMark indicates an expected call of DoUpdateMark.
func (mr *MockInterfaceDBMockRecorder) DoUpdateMark(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoUpdateMark", reflect.TypeOf((*MockInterfaceDB)(nil).DoUpdateMark), arg0, arg1)
}

// IsMyMark mocks base method.
func (m *MockInterfaceDB) IsMyMark(arg0 string, arg1 *[]Models.SubjectMarks) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMyMark", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsMyMark indicates an expected call of IsMyMark.
func (mr *MockInterfaceDBMockRecorder) IsMyMark(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMyMark", reflect.TypeOf((*MockInterfaceDB)(nil).IsMyMark), arg0, arg1)
}

// IsPresent mocks base method.
func (m *MockInterfaceDB) IsPresent(arg0 string, arg1 *Models.Student) (*Models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPresent", arg0, arg1)
	ret0, _ := ret[0].(*Models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPresent indicates an expected call of IsPresent.
func (mr *MockInterfaceDBMockRecorder) IsPresent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPresent", reflect.TypeOf((*MockInterfaceDB)(nil).IsPresent), arg0, arg1)
}

// IsPresentMark mocks base method.
func (m *MockInterfaceDB) IsPresentMark(arg0 string, arg1 *Models.SubjectMarks) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPresentMark", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsPresentMark indicates an expected call of IsPresentMark.
func (mr *MockInterfaceDBMockRecorder) IsPresentMark(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPresentMark", reflect.TypeOf((*MockInterfaceDB)(nil).IsPresentMark), arg0, arg1)
}

// MyMarks mocks base method.
func (m *MockInterfaceDB) MyMarks(arg0 string, arg1 *[]Models.SubjectMarks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MyMarks", arg0, arg1)
}

// MyMarks indicates an expected call of MyMarks.
func (mr *MockInterfaceDBMockRecorder) MyMarks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MyMarks", reflect.TypeOf((*MockInterfaceDB)(nil).MyMarks), arg0, arg1)
}
