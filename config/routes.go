package config

import (
	"github.com/darkphotonKN/go-figma-mcp/internal/figma"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up API routes and all routers
func SetupRouter(appConfig *AppConfig) *gin.Engine {
	router := gin.Default()

	// API base route
	api := router.Group("/api")

	// --- FIGMA ---

	// -- Figma Setup --
	figmaClient := figma.NewClient(appConfig.FigmaKey)
	figmaService := figma.NewService(figmaClient)
	figmaHandler := figma.NewHandler(figmaService)

	// -- Figma Routes --
	figmaRoutes := api.Group("/figma")
	figmaRoutes.GET("/files/:id", figmaHandler.GetFileInfo)

	return router
}

