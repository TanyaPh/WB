package controller

import (
	"api/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(services *service.Service) *gin.Engine {
	router := gin.New()

	router.Static("/orders", "./ui")
	newOrderRoutes(router.Group("/api"), services.Order)

	return router
}
