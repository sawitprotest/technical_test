package config

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

const (
	privateKeyPath = "./config/app.rsa"
	publicKeyPath  = "./config/app.rsa.pub"
)

func InitPublicKey() (*rsa.PublicKey, error) {
	verifyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}

func InitPrivateKey() (*rsa.PrivateKey, error) {
	signBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}
	return signKey, nil
}
