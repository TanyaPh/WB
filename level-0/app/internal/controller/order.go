package controller

import (
	"api/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type orderHandler struct {
	orderService service.Order
}

func newOrderRoutes(g *gin.RouterGroup, orderService service.Order) *orderHandler {
	h := &orderHandler{
		orderService: orderService,
	}

	g.GET("/orders/:id", h.getOrderByID)

	return h
}

func (h *orderHandler) getOrderByID(c *gin.Context) {
	idStr := c.Param("id")

	order, err := h.orderService.GetById(idStr)
	if err != nil {
		logrus.Errorf("No order with id: %s", idStr)
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("No order with id: %s", idStr))
	}

	c.JSON(http.StatusOK, order)
}
