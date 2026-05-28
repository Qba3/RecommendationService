package handler

import (
	"RecommendationService/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) GetProducts(
	w http.ResponseWriter,
	r *http.Request,
) {
	enableCors(w)

	log.Println("GET /products")

	products, err := h.service.GetAllProducts()
	if err != nil {
		log.Println("get products error:", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("products fetched:", len(products))

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println("products encode error:", err)
		return
	}

	log.Println("products response sent")
}
