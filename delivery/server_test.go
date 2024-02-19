package delivery

import (
	managermock "4crypto/mock/manager_mock"
	servicemock "4crypto/mock/service_mock"
	usecasemock "4crypto/mock/usecase_mock"

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
}

func (suite *ServerTestSuite) SetupTest() {
	suite.aucm = new(usecasemock.AuthUseCaseMock)
	suite.jtm = new(servicemock.JwtTokenMock)
	suite.umm = new(managermock.UseCaseManagerMock)
	suite.engine = gin.New()
	suite.host = ":8080"
}

func (suite *ServerTestSuite) TestNewServer_Success() {

}

func (suite *ServerTestSuite) TestNewServer_Failure() {

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
