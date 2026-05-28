package service

import (
	"RecommendationService/internal/model"
	"RecommendationService/internal/repository"
	"sort"
)

type RecommendationService struct {
	orderRepo   *repository.OrderRepository
	productRepo *repository.ProductRepository
}

func NewRecommendationService(
	orderRepo *repository.OrderRepository,
	productRepo *repository.ProductRepository,
) *RecommendationService {

	return &RecommendationService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *RecommendationService) GetRecommendations(
	userID int,
) ([]model.Product, error) {

	println("========== RECOMMENDATION START ==========")
	println("user id:", userID)

	userProducts, err := s.orderRepo.GetUserPurchasedProducts(userID)
	if err != nil {
		return nil, err
	}

	println("user purchased products:")

	for _, productID := range userProducts {
		println("-", productID)
	}

	if len(userProducts) == 0 {

		println("no purchases found")
		println("returning popular products")

		return s.productRepo.GetPopularProducts()
	}

	allOrders, err := s.orderRepo.GetAllUserProducts()
	if err != nil {
		return nil, err
	}

	println("loaded all user-product relations")

	matrix := buildProductMatrix(allOrders)

	println("matrix built")

	recommendationScores := map[int]int{}

	for _, productID := range userProducts {

		println("checking related products for:", productID)

		relatedProducts := matrix[productID]

		for relatedID, score := range relatedProducts {

			println(
				"related product:",
				relatedID,
				"score:",
				score,
			)

			if contains(userProducts, relatedID) {

				println(
					"already purchased, skipping:",
					relatedID,
				)

				continue
			}

			recommendationScores[relatedID] += score

			println(
				"updated recommendation score:",
				relatedID,
				"->",
				recommendationScores[relatedID],
			)
		}
	}

	type pair struct {
		ProductID int
		Score     int
	}

	var sorted []pair

	for productID, score := range recommendationScores {

		println(
			"final score:",
			productID,
			"=",
			score,
		)

		sorted = append(sorted, pair{
			ProductID: productID,
			Score:     score,
		})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Score > sorted[j].Score
	})

	println("sorted recommendations:")

	var recommendedIDs []int

	for _, item := range sorted {

		println(
			"recommended product:",
			item.ProductID,
			"score:",
			item.Score,
		)

		recommendedIDs = append(
			recommendedIDs,
			item.ProductID,
		)
	}

	if len(recommendedIDs) == 0 {

		println("no recommendations generated")
		println("returning popular products")

		return s.productRepo.GetPopularProducts()
	}

	println("========== RECOMMENDATION END ==========")

	return s.productRepo.GetProductsByIDs(recommendedIDs)
}

func buildProductMatrix(
	userOrders map[int][]int,
) map[int]map[int]int {

	println("========== BUILD MATRIX ==========")

	matrix := map[int]map[int]int{}

	for userID, products := range userOrders {

		println("processing user:", userID)

		for _, productID := range products {
			println("product:", productID)
		}

		for i := 0; i < len(products); i++ {

			for j := 0; j < len(products); j++ {

				if i == j {
					continue
				}

				productA := products[i]
				productB := products[j]

				if matrix[productA] == nil {
					matrix[productA] = map[int]int{}
				}

				matrix[productA][productB]++

				println(
					"matrix update:",
					productA,
					"->",
					productB,
					"score:",
					matrix[productA][productB],
				)
			}
		}
	}

	println("========== MATRIX READY ==========")

	return matrix
}

func contains(
	products []int,
	productID int,
) bool {

	for _, id := range products {

		if id == productID {
			return true
		}
	}

	return false
}
