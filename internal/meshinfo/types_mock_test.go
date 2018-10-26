// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qmuntal/go3mf/internal/meshinfo (interfaces: Invalidator,Container,MeshInfo)

// Package meshinfo is a generated GoMock package.
package meshinfo

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInvalidator is a mock of Invalidator interface
type MockInvalidator struct {
	ctrl     *gomock.Controller
	recorder *MockInvalidatorMockRecorder
}

// MockInvalidatorMockRecorder is the mock recorder for MockInvalidator
type MockInvalidatorMockRecorder struct {
	mock *MockInvalidator
}

// NewMockInvalidator creates a new mock instance
func NewMockInvalidator(ctrl *gomock.Controller) *MockInvalidator {
	mock := &MockInvalidator{ctrl: ctrl}
	mock.recorder = &MockInvalidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInvalidator) EXPECT() *MockInvalidatorMockRecorder {
	return m.recorder
}

// Invalidate mocks base method
func (m *MockInvalidator) Invalidate(arg0 interface{}) {
	m.ctrl.Call(m, "Invalidate", arg0)
}

// Invalidate indicates an expected call of Invalidate
func (mr *MockInvalidatorMockRecorder) Invalidate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockInvalidator)(nil).Invalidate), arg0)
}

// MockContainer is a mock of Container interface
type MockContainer struct {
	ctrl     *gomock.Controller
	recorder *MockContainerMockRecorder
}

// MockContainerMockRecorder is the mock recorder for MockContainer
type MockContainerMockRecorder struct {
	mock *MockContainer
}

// NewMockContainer creates a new mock instance
func NewMockContainer(ctrl *gomock.Controller) *MockContainer {
	mock := &MockContainer{ctrl: ctrl}
	mock.recorder = &MockContainerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContainer) EXPECT() *MockContainerMockRecorder {
	return m.recorder
}

// AddFaceData mocks base method
func (m *MockContainer) AddFaceData(arg0 uint32) (interface{}, error) {
	ret := m.ctrl.Call(m, "AddFaceData", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFaceData indicates an expected call of AddFaceData
func (mr *MockContainerMockRecorder) AddFaceData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFaceData", reflect.TypeOf((*MockContainer)(nil).AddFaceData), arg0)
}

// Clear mocks base method
func (m *MockContainer) Clear() {
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear
func (mr *MockContainerMockRecorder) Clear() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockContainer)(nil).Clear))
}

// Clone mocks base method
func (m *MockContainer) Clone() Container {
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(Container)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockContainerMockRecorder) Clone() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockContainer)(nil).Clone))
}

// GetCurrentFaceCount mocks base method
func (m *MockContainer) GetCurrentFaceCount() uint32 {
	ret := m.ctrl.Call(m, "GetCurrentFaceCount")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetCurrentFaceCount indicates an expected call of GetCurrentFaceCount
func (mr *MockContainerMockRecorder) GetCurrentFaceCount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentFaceCount", reflect.TypeOf((*MockContainer)(nil).GetCurrentFaceCount))
}

// GetFaceData mocks base method
func (m *MockContainer) GetFaceData(arg0 uint32) (interface{}, error) {
	ret := m.ctrl.Call(m, "GetFaceData", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFaceData indicates an expected call of GetFaceData
func (mr *MockContainerMockRecorder) GetFaceData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFaceData", reflect.TypeOf((*MockContainer)(nil).GetFaceData), arg0)
}

// MockMeshInfo is a mock of MeshInfo interface
type MockMeshInfo struct {
	ctrl     *gomock.Controller
	recorder *MockMeshInfoMockRecorder
}

// MockMeshInfoMockRecorder is the mock recorder for MockMeshInfo
type MockMeshInfoMockRecorder struct {
	mock *MockMeshInfo
}

// NewMockMeshInfo creates a new mock instance
func NewMockMeshInfo(ctrl *gomock.Controller) *MockMeshInfo {
	mock := &MockMeshInfo{ctrl: ctrl}
	mock.recorder = &MockMeshInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshInfo) EXPECT() *MockMeshInfoMockRecorder {
	return m.recorder
}

// AddFaceData mocks base method
func (m *MockMeshInfo) AddFaceData(arg0 uint32) (interface{}, error) {
	ret := m.ctrl.Call(m, "AddFaceData", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFaceData indicates an expected call of AddFaceData
func (mr *MockMeshInfoMockRecorder) AddFaceData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFaceData", reflect.TypeOf((*MockMeshInfo)(nil).AddFaceData), arg0)
}

// Clear mocks base method
func (m *MockMeshInfo) Clear() {
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear
func (mr *MockMeshInfoMockRecorder) Clear() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockMeshInfo)(nil).Clear))
}

// FaceHasData mocks base method
func (m *MockMeshInfo) FaceHasData(arg0 uint32) bool {
	ret := m.ctrl.Call(m, "FaceHasData", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// FaceHasData indicates an expected call of FaceHasData
func (mr *MockMeshInfoMockRecorder) FaceHasData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FaceHasData", reflect.TypeOf((*MockMeshInfo)(nil).FaceHasData), arg0)
}

// GetCurrentFaceCount mocks base method
func (m *MockMeshInfo) GetCurrentFaceCount() uint32 {
	ret := m.ctrl.Call(m, "GetCurrentFaceCount")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetCurrentFaceCount indicates an expected call of GetCurrentFaceCount
func (mr *MockMeshInfoMockRecorder) GetCurrentFaceCount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentFaceCount", reflect.TypeOf((*MockMeshInfo)(nil).GetCurrentFaceCount))
}

// GetFaceData mocks base method
func (m *MockMeshInfo) GetFaceData(arg0 uint32) (interface{}, error) {
	ret := m.ctrl.Call(m, "GetFaceData", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFaceData indicates an expected call of GetFaceData
func (mr *MockMeshInfoMockRecorder) GetFaceData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFaceData", reflect.TypeOf((*MockMeshInfo)(nil).GetFaceData), arg0)
}

// GetType mocks base method
func (m *MockMeshInfo) GetType() InformationType {
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(InformationType)
	return ret0
}

// GetType indicates an expected call of GetType
func (mr *MockMeshInfoMockRecorder) GetType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockMeshInfo)(nil).GetType))
}

// Invalidate mocks base method
func (m *MockMeshInfo) Invalidate(arg0 interface{}) {
	m.ctrl.Call(m, "Invalidate", arg0)
}

// Invalidate indicates an expected call of Invalidate
func (mr *MockMeshInfoMockRecorder) Invalidate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockMeshInfo)(nil).Invalidate), arg0)
}

// ResetFaceInformation mocks base method
func (m *MockMeshInfo) ResetFaceInformation(arg0 uint32) {
	m.ctrl.Call(m, "ResetFaceInformation", arg0)
}

// ResetFaceInformation indicates an expected call of ResetFaceInformation
func (mr *MockMeshInfoMockRecorder) ResetFaceInformation(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetFaceInformation", reflect.TypeOf((*MockMeshInfo)(nil).ResetFaceInformation), arg0)
}

// cloneFaceInfosFrom mocks base method
func (m *MockMeshInfo) cloneFaceInfosFrom(arg0 uint32, arg1 MeshInfo, arg2 uint32) {
	m.ctrl.Call(m, "cloneFaceInfosFrom", arg0, arg1, arg2)
}

// cloneFaceInfosFrom indicates an expected call of cloneFaceInfosFrom
func (mr *MockMeshInfoMockRecorder) cloneFaceInfosFrom(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cloneFaceInfosFrom", reflect.TypeOf((*MockMeshInfo)(nil).cloneFaceInfosFrom), arg0, arg1, arg2)
}

// getInternalID mocks base method
func (m *MockMeshInfo) getInternalID() uint64 {
	ret := m.ctrl.Call(m, "getInternalID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// getInternalID indicates an expected call of getInternalID
func (mr *MockMeshInfoMockRecorder) getInternalID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getInternalID", reflect.TypeOf((*MockMeshInfo)(nil).getInternalID))
}

// mergeInformationFrom mocks base method
func (m *MockMeshInfo) mergeInformationFrom(arg0 MeshInfo) {
	m.ctrl.Call(m, "mergeInformationFrom", arg0)
}

// mergeInformationFrom indicates an expected call of mergeInformationFrom
func (mr *MockMeshInfoMockRecorder) mergeInformationFrom(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mergeInformationFrom", reflect.TypeOf((*MockMeshInfo)(nil).mergeInformationFrom), arg0)
}

// permuteNodeInformation mocks base method
func (m *MockMeshInfo) permuteNodeInformation(arg0, arg1, arg2, arg3 uint32) {
	m.ctrl.Call(m, "permuteNodeInformation", arg0, arg1, arg2, arg3)
}

// permuteNodeInformation indicates an expected call of permuteNodeInformation
func (mr *MockMeshInfoMockRecorder) permuteNodeInformation(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "permuteNodeInformation", reflect.TypeOf((*MockMeshInfo)(nil).permuteNodeInformation), arg0, arg1, arg2, arg3)
}

// setInternalID mocks base method
func (m *MockMeshInfo) setInternalID(arg0 uint64) {
	m.ctrl.Call(m, "setInternalID", arg0)
}

// setInternalID indicates an expected call of setInternalID
func (mr *MockMeshInfoMockRecorder) setInternalID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setInternalID", reflect.TypeOf((*MockMeshInfo)(nil).setInternalID), arg0)
}
