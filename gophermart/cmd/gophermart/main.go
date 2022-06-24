package main

import (
	"fmt"
	"github.com/AXlIS/gofermart/internal/config"
	"github.com/AXlIS/gofermart/internal/handler"
	"github.com/AXlIS/gofermart/internal/server"
	s "github.com/AXlIS/gofermart/internal/service"
	store "github.com/AXlIS/gofermart/internal/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cfg, err := config.Init()
	if err != nil {
		log.Debug().Msgf("error: %s", err.Error())
	}

	db, err := store.NewPostgresDB(cfg.DB.DSN)
	if err != nil {
		log.Debug().Msgf("error: %s", err.Error())
	}

	storage := store.NewStorage(db)
	service := s.NewService(storage)
	router := handler.NewHandler(service)

	serve := new(server.Server)
	if err := serve.Start(fmt.Sprintf(":%s", cfg.HTTP.Port), router.InitRoutes()); err != nil {
		log.Debug().Msgf("error: %s", err.Error())
	}

	fmt.Println(cfg)
}
