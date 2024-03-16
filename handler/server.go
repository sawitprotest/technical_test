package handler

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawitpro/technical_test/config"
	"github.com/sawitpro/technical_test/repository"
	"github.com/sawitpro/technical_test/usecase"
)

type Server struct {
	cfg        *config.Config
	httpServer *echo.Echo
}

func InitServer() *Server {
	echoServer := echo.New()

	cfg := config.NewConfig()

	return &Server{
		cfg:        cfg,
		httpServer: echoServer,
	}
}

func (s *Server) Run() {

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Panic(err)
	}

	s.RegisterHandler()

	if err := s.httpServer.Start(fmt.Sprintf(":%d", serverPort)); err != nil {
		log.Println(err)
	}
}

func (s *Server) RegisterHandler() {

	repository := repository.NewRepository(s.cfg)
	userUsecase := usecase.NewUserUsecase(s.cfg, repository)
	userHandler := NewUserHandler(userUsecase)

	s.httpServer.POST(`/registration`, userHandler.Register)
	s.httpServer.POST(`/login`, userHandler.Login)
	s.httpServer.GET(`/profile/:id`, userHandler.GetProfile, JWTVerify(s.cfg.PublicKey))
	s.httpServer.PUT(`/profile/:id`, userHandler.UpdateProfile, JWTVerify(s.cfg.PublicKey))
}
