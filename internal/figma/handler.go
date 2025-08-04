package figma

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service     HandlerService
	figmaClient FigmaClient
}

type HandlerService interface {
	GetFileInfo(ctx context.Context, fileID string) error
}

type FigmaClient interface {
	GetFileInfo(fileID string) error
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetFileInfo(c *gin.Context) {
	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file ID is required"})
		return
	}

	err := h.figmaClient.GetFileInfo(fileID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file ID could not be retrieved, with error: " + err.Error()})
		return
	}

	err = h.service.GetFileInfo(c.Request.Context(), fileID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File info retrieved", "file_id": fileID})
}
