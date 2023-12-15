package repository_mock

import (
	"gin_app/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

type MockUserInformationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserInformationRepositoryMockRecorder
}

type MockUserInformationRepositoryMockRecorder struct {
	mock *MockUserInformationRepository
}

func NewMockUserInformationRepository(ctrl *gomock.Controller) *MockUserInformationRepository {
	mock := &MockUserInformationRepository{ctrl: ctrl}
	mock.recorder = &MockUserInformationRepositoryMockRecorder{mock}
	return mock
}

func (m *MockUserInformationRepository) EXPECT() *MockUserInformationRepositoryMockRecorder {
	return m.recorder
}

func (m *MockUserInformationRepository) GetByUserInformation(arg0 string) (*model.UserInformation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserInformation", arg0)
	ret0, _ := ret[0].(*model.UserInformation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserInformationRepositoryMockRecorder) GetByUserInformation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockUserInformationRepository)(nil).GetByUserInformation), arg0)
}
