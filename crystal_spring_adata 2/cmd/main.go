package main

import (
	"get_adata/internal/controller"
	"get_adata/internal/repository"
	"get_adata/internal/usecase"
	"log"
	"net/http"
)

func main() {
	baseURL := "http://10.1.22.179:8070"

	// Initialize the repository, use case, and controller
	tokenRepo := repository.NewTokenRepository(baseURL)
	companyUseCase := usecase.NewCompanyUseCase(tokenRepo)
	companyController := controller.NewCompanyController(companyUseCase)

	// Set up the HTTP route and handler
	http.HandleFunc("/company", companyController.GetCompanyData)

	// Start the HTTP server
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
