package delivery

import (
	"4crypto/config"

	// "bytes"
	// "errors"
	// "log"
	// "net/http"
	// "net/http/httptest"

	// "4crypto/delivery"
	configmock "4crypto/mock/config_mock"
	managermock "4crypto/mock/manager_mock"
	servicemock "4crypto/mock/service_mock"
	usecasemock "4crypto/mock/usecase_mock"
	"os"
	"time"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

var (
	apiConfigMock = config.ApiConfig{
		ApiPort: "8080",
	}
	dbConfigMock = config.DbConfig{
		Host:     "localhost",
		Port:     "5433",
		Name:     "enigma_laundry_apps",
		User:     "postgres",
		Password: "postgres",
		Driver:   "postgres",
	}

	tokenConfigMock = config.TokenConfig{
		IssuerName:      os.Getenv("TOKEN_ISSUER_NAME"),
		JwtSignatureKey: []byte(os.Getenv("TOKEN_KEY")),
		JwtLifeTime:     time.Duration(1) * time.Minute,
	}

	configMock = config.Config{
		ApiConfig:   apiConfigMock,
		DbConfig:    dbConfigMock,
		TokenConfig: tokenConfigMock,
	}
)

type ServerTestSuite struct {
	suite.Suite
	aucm   *usecasemock.AuthUseCaseMock
	jtm    *servicemock.JwtTokenMock
	engine *gin.Engine
	umm    *managermock.UseCaseManagerMock
	host   string
	cm     *configmock.ConfigMock
	// sm     *servermock.ServerMock
}

func (suite *ServerTestSuite) SetupTest() {
	suite.aucm = new(usecasemock.AuthUseCaseMock)
	suite.jtm = new(servicemock.JwtTokenMock)
	suite.umm = new(managermock.UseCaseManagerMock)
	suite.cm = new(configmock.ConfigMock)
	suite.engine = gin.Default()

	suite.host = ":8080"
}

func (suite *ServerTestSuite) TestSetupControllers_Success() {
	server := Server{
		engine:     suite.engine,
		auth:       suite.aucm,
		jwtService: suite.jtm,
	}

	server.setupControllers()
}

func (suite *ServerTestSuite) TestRun_Success() {

	NewServer()
}

func TestServerMockTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
