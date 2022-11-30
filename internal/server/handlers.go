package server

import (
	"final-project-backend/internal/admin/delivery"
	"final-project-backend/internal/admin/repository"
	"final-project-backend/internal/admin/usecase"
	authDelivery "final-project-backend/internal/auth/delivery"
	authRepository "final-project-backend/internal/auth/repository"
	authUseCase "final-project-backend/internal/auth/usecase"
	"final-project-backend/internal/middleware"
	userDelivery "final-project-backend/internal/user/delivery"
	userRepository "final-project-backend/internal/user/repository"
	userUseCase "final-project-backend/internal/user/usecase"
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

	userRepo := userRepository.NewUserRepository(s.db)
	userUC := userUseCase.NewUserUseCase(s.cfg, userRepo)
	userHandlers := userDelivery.NewUserHandlers(s.cfg, userUC, s.logger)

	adminRepo := repository.NewAdminRepository(s.db)
	adminUC := usecase.NewAdminUseCase(s.cfg, adminRepo)
	adminHandlers := delivery.NewAdminHandlers(s.cfg, adminUC, s.logger)

	mw := middleware.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)
	s.gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
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
	adminGroup := v1.Group("/admin")

	authDelivery.MapAuthRoutes(authGroup, authHandlers, mw)
	userDelivery.MapUserRoutes(userGroup, userHandlers, mw)
	delivery.MapAdminRoutes(adminGroup, adminHandlers, mw)

	return nil
}
