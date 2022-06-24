package config

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"time"
)

const (
	defaultAccessTokenTTL  = 5 * time.Minute
	defaultRefreshTokenTTL = 24 * time.Hour * 30
	defaultPasswordSalt    = "superSecretSalt"
	defaultSignedKey       = "superSecretSignedKey"
)

var (
	addr string
	uri  string
)

type (
	Config struct {
		Auth AuthConfig
		DB   DatabaseConfig
		HTTP HTTPConfig
	}

	AuthConfig struct {
		JWT          JWTConfig
		PasswordSalt string
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration
		RefreshTokenTTL time.Duration
		SigningKey      string
	}

	DatabaseConfig struct {
		DSN string
	}

	HTTPConfig struct {
		Port string
		Host string
	}
)

func init() {
	flag.StringVar(&addr, "a", "", "address to listen")
	flag.StringVar(&uri, "d", "", "database uri")
	flag.Parse()
}

func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	jwt := JWTConfig{
		AccessTokenTTL:  defaultAccessTokenTTL,
		RefreshTokenTTL: defaultRefreshTokenTTL,
		SigningKey:      GetEnv("JWT_SIGNING_KEY", defaultSignedKey),
	}

	db := DatabaseConfig{
		DSN: GetEnv("DATABASE_URI", uri),
	}

	address := strings.Split(GetEnv("RUN_ADDRESS", addr), ":")
	http := HTTPConfig{
		Host: address[0],
		Port: address[1],
	}

	cfg := Config{
		Auth: AuthConfig{
			JWT:          jwt,
			PasswordSalt: GetEnv("PASSWORD_SALT", defaultPasswordSalt),
		},
		DB:   db,
		HTTP: http,
	}

	return &cfg, nil
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
