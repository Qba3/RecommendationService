package handler

import (
	"RecommendationService/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(
	w http.ResponseWriter,
	r *http.Request,
) {
	enableCors(w)

	log.Println("POST /order")

	if r.Method == http.MethodOptions {
		log.Println("OPTIONS /order")
		return
	}

	var input service.CreateOrderInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println("decode order error:", err)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("creating order for user:", input.UserID)
	log.Println("items count:", len(input.Items))

	orderID, err := h.service.CreateOrder(input)
	if err != nil {
		log.Println("create order error:", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("order created:", orderID)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		log.Println("order response encode error:", err)
		return
	}

	log.Println("order response sent")
}
