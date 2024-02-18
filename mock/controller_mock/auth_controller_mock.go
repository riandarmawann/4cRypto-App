package controllermock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type AuthControllerMock struct {
	mock.Mock
}

func (a *AuthControllerMock) Route() {
	//a.Called()
}

func (a *AuthControllerMock) loginHandler(ctx *gin.Context) {
	a.Called(ctx)
}

func (a *AuthControllerMock) refreshTokenHandler(ctx *gin.Context) {
	a.Called(ctx)
}
