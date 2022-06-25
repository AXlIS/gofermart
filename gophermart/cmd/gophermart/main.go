package main

import (
	"fmt"
	"github.com/AXlIS/gofermart/internal/config"
	"github.com/AXlIS/gofermart/internal/handler"
	"github.com/AXlIS/gofermart/internal/server"
	s "github.com/AXlIS/gofermart/internal/service"
	store "github.com/AXlIS/gofermart/internal/storage"
	"github.com/AXlIS/gofermart/pkg/auth"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cfg, err := config.Init()
	if err != nil {
		log.Fatal().Msgf("config error: %s", err.Error())
	}

	db, err := store.NewPostgresDB(cfg.DB.DSN)
	if err != nil {
		log.Fatal().Msgf("postgres error: %s", err.Error())
	}
	tokenManager, err := auth.NewManager(cfg.Auth.JWT)
	if err != nil {
		log.Fatal().Msgf("token manager error: %s", err.Error())
	}

	storage := store.NewStorage(db)
	service := s.NewService(storage, tokenManager)
	router := handler.NewHandler(service, tokenManager, cfg.Auth.JWT.AccessTokenTTL)

	serve := new(server.Server)
	if err := serve.Start(fmt.Sprintf(":%s", cfg.HTTP.Port), router.InitRoutes()); err != nil {
		log.Fatal().Msgf("server error: %s", err.Error())
	}
}
