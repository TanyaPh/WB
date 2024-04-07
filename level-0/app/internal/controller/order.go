package controller

import (
	"api/internal/service"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService service.Order
}

func newOrderRoutes(g *gin.RouterGroup, orderService service.Order) *orderHandler {
	h := &orderHandler{
		orderService: orderService,
	}

	g.GET("/:id", h.getOrderByID)

	return h
}

func (h *orderHandler) getOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	h.orderService.GetById(idStr)
}
