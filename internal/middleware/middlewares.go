package middleware

import (
	"final-project-backend/config"
	"final-project-backend/pkg/logger"
)

type MWManager struct {
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

func NewMiddlewareManager(cfg *config.Config, origins []string, logger logger.Logger) *MWManager {
	return &MWManager{cfg: cfg, origins: origins, logger: logger}
}
