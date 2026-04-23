package handler

import (
	"RecommendationService/internal/service"
	"net/http"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {}
