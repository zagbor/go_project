package handlers

import (
	"encoding/json"
	"go-microservice/services"
	"net/http"
)

// IntegrationHandler handles requests related to external integrations.
type IntegrationHandler struct {
	Service *services.IntegrationService
}

// NewIntegrationHandler creates a new handler.
func NewIntegrationHandler(s *services.IntegrationService) *IntegrationHandler {
	return &IntegrationHandler{Service: s}
}

// HealthCheck is a simple endpoint to verify integration connectivity.
func (h *IntegrationHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{"status": "Integration Service Operational"}
	
	// In a real scenario, we might check MinIO connectivity here
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
