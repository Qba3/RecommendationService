package main

import (
	"RecommendationService/internal/db"
	"RecommendationService/internal/handler"
	"RecommendationService/internal/repository"
	"RecommendationService/internal/service"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	database := db.NewDB()

	orderRepo := repository.NewOrderRepository(database)
	orderItemRepo := repository.NewOrderItemRepository(database)
	productRepo := repository.NewProductRepository(database)
	userRepo := repository.NewUserRepository(database)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	recommendationService := service.NewRecommendationService(
		orderRepo,
		productRepo,
	)

	orderService := service.NewOrderService(
		orderRepo,
		orderItemRepo,
		productRepo,
		recommendationService,
	)

	orderHandler := handler.NewOrderHandler(orderService)

	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	recommendationHandler := handler.NewRecommendationHandler(recommendationService)

	http.HandleFunc(
		"/recommendations/",
		recommendationHandler.GetRecommendations,
	)

	http.HandleFunc("/login", userHandler.Login)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/order", orderHandler.CreateOrder)
	http.HandleFunc("/products", productHandler.GetProducts)

	log.Println("server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
