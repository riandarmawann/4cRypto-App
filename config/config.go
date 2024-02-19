package config

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	// "os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
}

type TokenConfig struct {
	IssuerName      string
	JwtSignatureKey []byte
	JwtLifeTime     time.Duration
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func (c *Config) readConfig() error {

	if err := godotenv.Load(basepath + "/../.env"); err != nil {
		return errors.New("failed to load environment variables")
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))
	if err != nil {
		return err
	}

	c.TokenConfig = TokenConfig{
		IssuerName:      os.Getenv("TOKEN_ISSUER_NAME"),
		JwtSignatureKey: []byte(os.Getenv("TOKEN_KEY")),
		JwtLifeTime:     time.Duration(tokenLifeTime) * time.Minute,
	}

	if c.ApiPort == "" || c.Host == "" || c.Port == "" || c.Name == "" || c.User == "" || c.IssuerName == "" || c.JwtSignatureKey == nil || c.JwtLifeTime == 0 {
		return errors.New("environment required")
	}

	return nil
}
