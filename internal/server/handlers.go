package server

import (
	authDelivery "final-project-backend/internal/auth/delivery"
	authRepository "final-project-backend/internal/auth/repository"
	authUseCase "final-project-backend/internal/auth/usecase"
	"final-project-backend/internal/middleware"
	"final-project-backend/internal/user/delivery"
	"final-project-backend/internal/user/repository"
	"final-project-backend/internal/user/usecase"
	"final-project-backend/pkg/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *Server) MapHandlers() error {
	aRepo := authRepository.NewAuthRepository(s.db)
	authUC := authUseCase.NewAuthUseCase(s.cfg, aRepo)
	authHandlers := authDelivery.NewAuthHandlers(s.cfg, authUC, s.logger)

	userRepo := repository.NewUserRepository(s.db)
	userUC := usecase.NewUserUseCase(s.cfg, userRepo)
	userHandlers := delivery.NewUserHandlers(s.cfg, userUC, s.logger)

	mw := middleware.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)
	s.gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3001"
		},
		MaxAge: 12 * time.Hour,
	}))

	s.gin.NoRoute(func(c *gin.Context) {
		response.ErrorResponse(c.Writer, response.NotFoundMessage, http.StatusNotFound)
	})

	v1 := s.gin.Group("/api/v1")
	authGroup := v1.Group("/auth")
	userGroup := v1.Group("/user")
	authDelivery.MapAuthRoutes(authGroup, authHandlers, mw)
	delivery.MapUserRoutes(userGroup, userHandlers, mw)

	return nil
}
