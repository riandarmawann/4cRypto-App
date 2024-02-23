package controllermock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type CryptoControllerMock struct {
	mock.Mock
}

func (a *CryptoControllerMock) Route() {
	//a.Called()
}

func (c *CryptoControllerMock) handleGetBook(ctx *gin.Context) {
	c.Called(ctx)
}
func (c *CryptoControllerMock) handleGetRank(ctx *gin.Context) {
	c.Called(ctx)
}

func (c *CryptoControllerMock) handlerPlaceOrder(ctx *gin.Context) {
	c.Called(ctx)
}

func (c *CryptoControllerMock) cancelOrder(ctx *gin.Context) {
	c.Called(ctx)
}
