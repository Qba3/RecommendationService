package service

import (
	"RecommendationService/internal/model"
	"RecommendationService/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.repo.GetAllProducts()
}
