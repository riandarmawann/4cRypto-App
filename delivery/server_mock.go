package delivery

import (
	"4crypto/usecase"
	"4crypto/utils/common"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type ServerMock struct {
	mock.Mock
	engine     *gin.Engine
	auth       usecase.AuthUseCase
	jwtService common.JwtToken
	host       string
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
