package controller

import (
	usecasemock "4crypto/mock/usecase_mock"
	cryptoEntity "4crypto/model/entity/crypto"
	"4crypto/utils/common"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type CryptoControllerTestSuite struct {
	suite.Suite
	cum usecasemock.CryptoUseCaseMock
	rg  *gin.RouterGroup

	cryptoEntity.Exchange
}

func (suite *CryptoControllerTestSuite) SetupTest() {
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1").Group("crypto")
	orderbooks := make(map[cryptoEntity.Market]*common.Orderbook)
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	orderbooks[cryptoEntity.MarketETH] = common.NewOrderBook()
	suite.Exchange = cryptoEntity.Exchange{
		Client:     client,
		Users:      make(map[int64]*cryptoEntity.User),
		Orders:     make(map[int64]int64),
		Orderbooks: orderbooks,
	}
}

func (suite *CryptoControllerTestSuite) TestRoute() {
	uc := NewCryptoController(&suite.cum, suite.rg, suite.Client)
	uc.Route()
	suite.rg.GET("/book/:market", uc.handleGetBook)
	suite.rg.GET("/rank", uc.handleGetRank)
	suite.rg.POST("/order", uc.handlerPlaceOrder)
	suite.rg.DELETE("/order/:id", uc.cancelOrder)
}

func (suite *CryptoControllerTestSuite) TestGetBook() {

	suite.cum.On("Orderbooks", common.Orderbook{}).Return(cryptoEntity.OrderbookData{})

	// Membuat request HTTP tes
	req, _ := http.NewRequest("GET", "/api/v1/crypto/book/:market", nil)
	w := httptest.NewRecorder()

	// Membuat context Gin untuk mensimulasikan request
	// c := new(gin.Context)
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "market", Value: "ETH"})

	uc := NewCryptoController(&suite.cum, suite.rg, suite.Client)
	// NewUserController(suite.ucm, suite.rg)
	uc.handleGetBook(c)

	suite.cum.AssertCalled(suite.T(), "Orderbooks", "1")
}
