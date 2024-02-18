package controllermock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type AuthControllerMock struct {
	mock.Mock
}

func (a *AuthControllerMock) Route() {
	//args := a.Called()
	//return args.Get(0).(entity.User), args.Error(1)
}

func (a *AuthControllerMock) loginHandler(ctx *gin.Context) {
	a.Called(ctx)
	//return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (a *AuthControllerMock) refreshTokenHandler(ctx *gin.Context) {
	a.Called(ctx)
	// return args.Get(0).(string), args.Error(1)
}
