package server

import (
	"final-project-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) MapHandlers() error {
	s.gin.NoRoute(func(c *gin.Context) {
		response.ErrorResponse(c.Writer, response.NotFoundMessage, http.StatusNotFound)
	})

	v1 := s.gin.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		response.ErrorResponse(c.Writer, response.VersionMessage, http.StatusOK)
	})

	return nil
}
