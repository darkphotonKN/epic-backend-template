package healthmetrics

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

type Service interface {
	ProcessGarminWebhook(ctx context.Context, request GarminWebhookPayload) (*HealthMetricProcessedData, error)
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// TODO: KIKIIIIIIIIIII
// processes incoming webhook data from Garmin Connect
func (h *Handler) HandleGarminWebhook(c *gin.Context) {
	var payload GarminWebhookPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Printf("Error parsing Garmin webhook payload: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Invalid webhook payload format",
			"result":     nil,
		})
		return
	}

	// Process the health metrics through our service
	processedData, err := h.Service.ProcessGarminWebhook(c.Request.Context(), payload)

	// handle service error response
	if err != nil {
		log.Printf("Error processing Garmin webhook: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    fmt.Sprintf("Failed to process health data: %s", err.Error()),
			"result":     nil,
		})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Health data processed successfully",
		"result":     processedData,
	})
}
