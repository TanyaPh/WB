package controller

import (
	"api/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(services *service.Service) *gin.Engine {
	router := gin.New()
	newOrderRoutes(router.Group("/orders"), services.Order)
	
	return router
}
