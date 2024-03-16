package config

import (
	"crypto/rsa"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DB         *gorm.DB
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

func NewConfig() *Config {
	privateKey, publicKey := InitKey()

	return &Config{
		DB:         InitDB(),
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
}

func InitKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	publicKey, err := InitPublicKey()
	if err != nil {
		log.Panic(err)
	}

	privateKey, err := InitPrivateKey()
	if err != nil {
		log.Panic(err)
	}

	return privateKey, publicKey
}

func InitDB() *gorm.DB {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	dbDsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=UTC+7",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	dbc, err := sql.Open("postgres", dbDsn)
	if err != nil {
		log.Panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbc,
	}), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		log.Panic(err)
	}

	return db
}
