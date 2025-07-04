package config

import (
	"github.com/darkphotonKN/eh-hub-data-orchestration-platform/internal/healthmetrics"
	"github.com/gin-gonic/gin"
)

/**
* Sets up API prefix route and all routers.
**/
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// base route
	api := router.Group("/api")

	// -- HEALTH METRICS --

	// --- Health Metrics Setup ---
	healthMetricsService := healthmetrics.NewService()
	healthMetricsHandler := healthmetrics.NewHandler(healthMetricsService)

	// --- Webhook Routes ---
	// Garmin webhook endpoint for receiving real-time health data
	webhookRoutes := api.Group("/webhooks")
	webhookRoutes.POST("/garmin", healthMetricsHandler.HandleGarminWebhook)

	return router
}