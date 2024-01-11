package usecase

import (
	"gin_app/domain/model"
	mock_repository "gin_app/domain/repository/repository_mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_userInformationUseCase_GetByUserInformation(t1 *testing.T) {
	type args struct {
		id string
	}

	wantUserInformation := model.UserInformation{
		Id:        1,
		LastName:  "山田",
		FirstName: "太郎",
		Sex:       "男",
		Age:       25,
		PostCode:  "111-1111",
		Address:   "東京都港区1-1-1",
		Remarks:   "",
	}

	tests := []struct {
		name          string
		uu            userInformationUseCase
		args          args
		prepareMockFn func(m *mock_repository.MockUserInformationRepository)
		wantUser      *model.UserInformation
		wantErr       bool
	}{
		{
			name: "success",
			args: args{id: "1"},
			prepareMockFn: func(m *mock_repository.MockUserInformationRepository) {
				m.EXPECT().GetByUserInformation("1").Return(&wantUserInformation, nil)
			},
			wantUser: &wantUserInformation,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {

			ctrl := gomock.NewController(t1)
			defer ctrl.Finish()

			mock := mock_repository.NewMockUserInformationRepository(ctrl)
			tt.prepareMockFn(mock)

			t := userInformationUseCase{
				userInformationRepository: mock,
			}
			gotUser, err := t.GetByUserInformation(tt.args.id)
			if (err != nil) != tt.wantErr {
				t1.Errorf("userInformationUseCase.GetByUserInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t1.Errorf("userInformationUseCase.GetByUserInformation() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
