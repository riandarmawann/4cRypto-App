package servermock

import (
	managermock "4crypto/mock/manager_mock"
	servicemock "4crypto/mock/service_mock"
	usecasemock "4crypto/mock/usecase_mock"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type ServerMock struct {
	mock.Mock
	engine         *gin.Engine
	host           string
	ucManagerMock  managermock.UseCaseManagerMock
	authMock       usecasemock.AuthUseCaseMock
	jwtServiceMock servicemock.JwtTokenMock
}

func (s *ServerMock) NewServer() *ServerMock {
	// Create mocked instances of dependencies
	engine := &gin.Engine{}
	ucManagerMock := &managermock.UseCaseManagerMock{}
	authMock := &usecasemock.AuthUseCaseMock{}
	jwtServiceMock := &servicemock.JwtTokenMock{}

	// Assuming host is ":8080" for simplicity
	host := ":8080"

	// Simulate the behavior of initializing a Server instance
	mockedServer := &ServerMock{
		engine:         engine,
		host:           host,
		ucManagerMock:  *ucManagerMock,
		authMock:       *authMock,
		jwtServiceMock: *jwtServiceMock,
	}

	return mockedServer
}

func (s *ServerMock) setupControllers() {
	rg := s.engine.Group("/api/v1")
	// Simulate the Route function call
	s.Called(rg)
}

func (s *ServerMock) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}
