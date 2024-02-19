package delivery

import (
	"4crypto/config"
	"4crypto/mock/server_mock"
	configmock "4crypto/mock/config_mock"
	managermock "4crypto/mock/manager_mock"
	servicemock "4crypto/mock/service_mock"
	usecasemock "4crypto/mock/usecase_mock"
	"errors"

	"bytes"

	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	aucm   *usecasemock.AuthUseCaseMock
	jtm    *servicemock.JwtTokenMock
	engine *gin.Engine
	umm    *managermock.UseCaseManagerMock
	host   string
	cm     *configmock.ConfigMock
	sm     *servermock.ServerMock
}

func (suite *ServerTestSuite) SetupTest() {
	suite.aucm = new(usecasemock.AuthUseCaseMock)
	suite.jtm = new(servicemock.JwtTokenMock)
	suite.umm = new(managermock.UseCaseManagerMock)
	suite.cm = new(configmock.ConfigMock)
	suite.engine = gin.New()
	suite.host = ":8080"
}

func (suite *ServerTestSuite) TestNewServer_Success() {
	// Mocks setup
	suite.cm.On("NewConfig").Return(&config.Config{}, nil)
	suite.umm.On("NewAuthUseCase").Return(suite.aucm)
	suite.umm.On("NewJwtTokenService").Return(suite.jtm)

	// Create a new server
	server, err := suite.cm.NewServer(suite.cm, suite.umm, suite.host)

	// Assertions
	assert.NotNil(suite.T(), server)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), server.UCManager)
	assert.NotNil(suite.T(), server.Engine)
	assert.NotNil(suite.T(), server.Auth)
	assert.NotNil(suite.T(), server.JwtService)
	assert.Equal(suite.T(), ":8080", server.Host)
}

func (suite *ServerTestSuite) TestNewServer_Failure() {
	// Mocks setup
	suite.cm.On("NewConfig").Return(nil, errors.New("configuration error"))

	// Attempt to create a new server with failing config
	server, err := suite.cm.NewServer(suite.cm, suite.umm, suite.host)

	// Assertions
	assert.Nil(suite.T(), server)
	assert.Error(suite.T(), err)
}

func (suite *ServerTestSuite) TestSetupControllers_Success() {
	// Persiapkan server
	server := &Server{
		engine: gin.New(),
	}

	// Jalankan setupControllers
	server.setupControllers()

	// pengujian untuk rute "/api/v1" yang diasumsikan memiliki controller AuthController
	req, err := http.NewRequest("GET", "/api/v1", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	rr := httptest.NewRecorder()
	server.engine.ServeHTTP(rr, req)

	// Lakukan pengujian untuk memastikan bahwa status code yang diharapkan diperoleh
	assert.Equal(suite.T(), http.StatusNotFound, rr.Code, "Handler returned wrong status code")
}

// func (suite *ServerTestSuite) TestRun() {
// 	// Persiapkan server
// 	server := &Server{
// 		engine: gin.New(),
// 		host:   ":8080",
// 	}

// 	// Jalankan setupControllers untuk menambahkan rute
// 	server.setupControllers()

// 	// Buat request palsu ke suatu rute yang telah ditetapkan
// 	req, err := http.NewRequest("GET", "/api/v1/route", nil)
// 	if err != nil {
// 		log.Fatalf("Error creating request: %v", err)
// 	}

// 	rr := httptest.NewRecorder()
// 	server.engine.ServeHTTP(rr, req)

// 	// Lakukan pengujian untuk memastikan bahwa respons status yang diharapkan diperoleh
// 	assert.Equal(suite.T(), http.StatusNotFound, rr.Code, "Handler returned wrong status code")
// }

func (suite *ServerTestSuite) TestRun() {
}

func (suite *ServerTestSuite) TestRun_Fail() {
	// Persiapkan server
	server := &Server{
		engine: gin.New(),
		host:   ":invalid", // Sengaja memasukkan host yang tidak valid untuk menimbulkan error
	}

	// Membuat buffer untuk menangkap output log
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Jalankan setupControllers untuk menambahkan rute
	server.setupControllers()

	// Jalankan fungsi Run yang seharusnya menghasilkan error
	server.Run()

	// Memeriksa apakah log.Fatal() dipanggil dengan pesan yang diharapkan
	assert.Contains(suite.T(), buf.String(), "server can't run", "log.Fatal() should be called with correct message")
}

func TestServerMockTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
