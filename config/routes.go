package config

import (
	"github.com/darkphotonKN/eh-hub-data-orchestration-platform/internal/item"
	"github.com/gin-gonic/gin"
)

/**
* Sets up API prefix route and all routers.
**/
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// base route
	api := router.Group("/api")

	// -- ITEM --

	// --- Item Setup ---
	itemService := item.NewService()
	itemHandler := item.NewHandler(itemService)

	// --- Item Routes ---
	itemRoutes := api.Group("/items")
	itemRoutes.POST("/", itemHandler.Create)
	itemRoutes.GET("/", itemHandler.GetAll)
	itemRoutes.GET("/:id", itemHandler.GetById)
	itemRoutes.PUT("/:id", itemHandler.Update)
	itemRoutes.DELETE("/:id", itemHandler.Delete)

	return router
}

