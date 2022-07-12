package main

import (
	"fmt"
	"github.com/AXlIS/gofermart/internal/config"
	"github.com/AXlIS/gofermart/internal/handler"
	"github.com/AXlIS/gofermart/internal/server"
	s "github.com/AXlIS/gofermart/internal/service"
	store "github.com/AXlIS/gofermart/internal/storage"
	"github.com/AXlIS/gofermart/pkg/auth"
	"github.com/AXlIS/gofermart/pkg/hash"
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

	hasher := hash.NewSHA256Hasher(cfg.Auth.PasswordSalt)

	storage := store.NewStorage(db)
	service := s.NewService(storage, tokenManager, hasher)
	router := handler.NewHandler(service, tokenManager, cfg.Auth.JWT.AccessTokenTTL)

	go func() {
		accrualServer := new(server.Server)
		accrualRouter := handler.NewAccrualHandler()
		if err := accrualServer.Start(fmt.Sprintf(":%s", cfg.HTTP.AccrualPort), accrualRouter.InitRoutes()); err != nil {
			log.Fatal().Msgf("server error: %s", err.Error())
		}
	}()

	serve := new(server.Server)
	if err := serve.Start(fmt.Sprintf(":%s", cfg.HTTP.Port), router.InitRoutes()); err != nil {
		log.Fatal().Msgf("server error: %s", err.Error())
	}
}
