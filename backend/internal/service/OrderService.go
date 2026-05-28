package service

import (
	"RecommendationService/internal/repository"
)

type OrderItemInput struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CreateOrderInput struct {
	UserID int              `json:"user_id"`
	Items  []OrderItemInput `json:"items"`
}

type OrderService struct {
	repo                  *repository.OrderRepository
	orderItemRepo         *repository.OrderItemRepository
	productRepo           *repository.ProductRepository
	recommendationService *RecommendationService
}

func NewOrderService(
	repo *repository.OrderRepository,
	orderItemRepo *repository.OrderItemRepository,
	productRepo *repository.ProductRepository,
	recommendationService *RecommendationService,
) *OrderService {
	return &OrderService{
		repo:                  repo,
		orderItemRepo:         orderItemRepo,
		productRepo:           productRepo,
		recommendationService: recommendationService,
	}
}

func (s *OrderService) CreateOrder(input CreateOrderInput) (int, error) {
	orderID, err := s.repo.CreateOrder(input.UserID)
	if err != nil {
		return 0, err
	}

	for _, item := range input.Items {
		err := s.orderItemRepo.CreateOrderItem(
			orderID,
			item.ProductID,
			item.Quantity,
		)

		if err != nil {
			return 0, err
		}
	}

	return orderID, nil
}
