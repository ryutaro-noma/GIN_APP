package usecase

import (
	"gin_app/domain/model"
	"gin_app/domain/repository/repository_mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func Test_GetByUserInformation(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := repository_mock.NewMockUserInformationRepository(mockCtrl)
	mock.EXPECT().GetByUserInformation("1").Return(
		&model.UserInformation{
			Id:        1,
			LastName:  "山田",
			FirstName: "太郎",
			Sex:       "男",
			Age:       25,
			PostCode:  "111-1111",
			Address:   "東京都港区1-1-1",
			Remarks:   "",
		},
	)

	var want *model.UserInformation = &model.UserInformation{
		Id:        1,
		LastName:  "山田",
		FirstName: "太郎",
		Sex:       "男",
		Age:       25,
		PostCode:  "111-1111",
		Address:   "東京都港区1-1-1",
		Remarks:   "",
	}

	got, err := UserInformationUseCase.GetByUserInformation(mock, "1")
	if err != nil {
		t.Error("error happen")
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("UserInformation Data miss match :%s", diff)
	}
}
