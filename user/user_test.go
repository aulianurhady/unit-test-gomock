package user_test

import (
	"testing"

	"github.com/authocode/sample/mocks"
	"github.com/authocode/sample/user"
	"github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDoer := mocks.NewMockDoer(mockController)
	testUser := &user.User{
		Doer: mockDoer,
	}

	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}
