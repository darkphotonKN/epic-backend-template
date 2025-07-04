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
	ProcessGarminWebhook(ctx context.Context, payload GarminWebhookPayload) (int, error)
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
		c.JSON(http.StatusBadRequest, WebhookResponse{
			Status:  "error",
			Message: "Invalid webhook payload format",
		})
		return
	}

	// Process the health metrics through our service
	processedCount, err := h.Service.ProcessGarminWebhook(c.Request.Context(), payload)

	// handle service error response
	if err != nil {
		log.Printf("Error processing Garmin webhook: %v", err)
		c.JSON(http.StatusInternalServerError, WebhookResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to process health data: %s", err.Error()),
		})
		return
	}

	// success response
	c.JSON(http.StatusOK, WebhookResponse{
		Status:    "success",
		Message:   "Health data processed successfully",
		Processed: processedCount,
	})
}

