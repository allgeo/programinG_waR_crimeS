package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHttpHandler(s *Service) {

	handler := &HTTPHandler{s}
	v1 := s.Router.Group("/api")
	{
		// Health Check
		v1.GET("/ping", HealthCheck)

		// Create and Get all Posts
		v1.POST("/post", handler.CreatePost)
		v1.GET("/post", handler.GetAllPost)
		v1.POST("/post/upload", handler.UploadPostFile)

	}
}

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
