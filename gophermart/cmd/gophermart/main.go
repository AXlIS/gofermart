package main

import (
	"fmt"
	"github.com/AXlIS/gofermart/internal/config"
	store "github.com/AXlIS/gofermart/internal/storage"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	db, err := store.NewPostgresDB(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}


	fmt.Println(cfg)
}
