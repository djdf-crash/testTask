package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"testTask/cmd/internal/config"
)

type IServer interface {
	Close() error
	Start() error
	InitRoutes(s IHandlerService)
	RegisterMiddleware(s IAuthMiddlewareService)
}

type Server struct {
	f   *fiber.App
	cfg config.ServerConfig
}

func (s *Server) Close() error {
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
	defer cancel()

	return s.f.ShutdownWithContext(ctx)
}

func (s *Server) Start() error {
	return s.f.Listen(":" + s.cfg.Port)
}

func (s *Server) InitRoutes(service IHandlerService) {
	h := NewHandler(service)
	s.f.Get("/profile", h.UserProfiles)
}

func (s *Server) RegisterMiddleware(service IAuthMiddlewareService) {
	s.f.Use(NewAuthMiddleware(service))
}

func NewServer(cfg config.ServerConfig) IServer {
	f := fiber.New()
	return &Server{f: f, cfg: cfg}
}
