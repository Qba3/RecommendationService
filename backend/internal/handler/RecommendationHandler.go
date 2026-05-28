package handler

import (
	"RecommendationService/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type RecommendationHandler struct {
	service *service.RecommendationService
}

func NewRecommendationHandler(
	service *service.RecommendationService,
) *RecommendationHandler {

	return &RecommendationHandler{
		service: service,
	}
}

func (h *RecommendationHandler) GetRecommendations(
	w http.ResponseWriter,
	r *http.Request,
) {
	enableCors(w)

	userIDStr := strings.TrimPrefix(
		r.URL.Path,
		"/recommendations/",
	)

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	products, err := h.service.GetRecommendations(userID)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(products)
}
