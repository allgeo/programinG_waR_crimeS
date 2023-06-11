package controllers

import (
	s3 "blog-apis/api/clients/s3service"
	weather "blog-apis/api/clients/weatherservice"
	"blog-apis/api/models"

	"github.com/gin-gonic/gin"
)

type Service struct {
	AwsS3Service   *s3.AwsS3ServiceClient
	WeatherService *weather.WeatherServiceClient
	Repo           models.DBRepositoryInterface
	Router         *gin.Engine
}

type HTTPHandler struct {
	service *Service
}

//NewFeedService creates and returns the feedservice
func NewService(sc *s3.AwsS3ServiceClient, wc *weather.WeatherServiceClient, r models.DBRepositoryInterface, router *gin.Engine) *Service {
	return &Service{sc, wc, r, router}
}
