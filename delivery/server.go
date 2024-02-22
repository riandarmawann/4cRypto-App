package delivery

import (
	"4crypto/usecase"
	"4crypto/utils/common"
	"fmt"
	"log"

	"4crypto/config"
	"4crypto/delivery/controller"
	"4crypto/delivery/middleware"
	"4crypto/manager"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ucManager  manager.UseCaseManager
	auth       usecase.AuthUseCase
	engine     *gin.Engine
	client     *ethclient.Client
	host       string
	jwtService common.JwtToken
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	repoManager := manager.NewRepoManager(infraManager)
	ucManager := manager.NewUseCaseManager(repoManager)
	engine := gin.New()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	jwtService := common.NewJwtToken(cfg.TokenConfig)

	return &Server{
		ucManager:  ucManager,
		engine:     engine,
		host:       host,
		client:     client,
		auth:       usecase.NewAuthUseCase(ucManager.NewUserUseCase(), jwtService),
		jwtService: jwtService,
	}
}

func (s *Server) setupControllers() {
	loggerMiddleware := middleware.NewLoggerMiddleware().Logger()
	rg := s.engine.Group("/api/v1", loggerMiddleware)
	controller.NewAuthController(s.auth, rg, s.jwtService).Route()
	controller.NewCryptoController(s.ucManager.NewCryptoUseCase(), rg, s.client).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}
