package figma

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for the domain
type Handler struct {
	service Service
}

// NewHandler creates a new handler instance
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// GetFileInfo handles GET /files/:id requests
func (h *Handler) GetFileInfo(c *gin.Context) {
	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file ID is required"})
		return
	}

	err := h.service.GetFileInfo(c.Request.Context(), fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File info retrieved", "file_id": fileID})
}
