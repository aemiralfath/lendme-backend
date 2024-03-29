package server

import (
	"context"
	"final-project-backend/config"
	"final-project-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Server struct {
	gin    *gin.Engine
	cfg    *config.Config
	db     *gorm.DB
	logger logger.Logger
}

func NewServer(cfg *config.Config, db *gorm.DB, logger logger.Logger) *Server {
	return &Server{gin: gin.Default(), cfg: cfg, db: db, logger: logger}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		Handler:        s.gin,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	s.logger.Info("Shutdown Server ...")

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
		return err
	}

	select {
	case <-ctx.Done():
		log.Println("timeout.")
	}
	s.logger.Info("Server Exited Properly")

	return nil
}
