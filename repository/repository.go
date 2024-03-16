package repository

import (
	"math/rand"
	"time"

	"github.com/sawitpro/technical_test/config"
)

type repositoryCtx struct {
	cfg *config.Config
}

func NewRepository(cfg *config.Config) Repository {
	rand.Seed(time.Now().UnixNano())
	return &repositoryCtx{
		cfg: cfg,
	}
}
