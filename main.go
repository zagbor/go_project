package main

import (
	"go-microservice/handlers"
	"go-microservice/metrics"
	"go-microservice/services"
	"go-microservice/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Initialize Metrics
	metrics.Init()

	// Initialize Services and Handlers
	userService := services.NewUserService()
	userHandler := handlers.NewUserHandler(userService)

	// Stub configuration for MinIO (Replace with real env vars in production)
	integrationService := services.NewIntegrationService("localhost:9000", "minioadmin", "minioadmin")
	integrationHandler := handlers.NewIntegrationHandler(integrationService)

	// Setup Router
	r := mux.NewRouter()

	// Middleware
	r.Use(utils.RateLimitMiddleware) // Applied to all routes
	r.Use(metrics.MetricsMiddleware) // Collections metrics for all routes

	// API Routes
	r.HandleFunc("/api/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")

	// Integration Routes
	r.HandleFunc("/api/integration/health", integrationHandler.HealthCheck).Methods("GET")

	// Prometheus Metrics Endpoint
	r.Handle("/metrics", promhttp.Handler())

	// Start Server
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
