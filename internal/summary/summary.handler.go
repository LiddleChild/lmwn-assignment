package summary

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetSummary(c *gin.Context)
}

type handlerImpl struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handlerImpl{
		service,
	}
}

func (h *handlerImpl) GetSummary(c *gin.Context) {
	var summary Summary
	err := h.service.GetSummary(&summary)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"name": err.Name,
			"path": c.Request.URL.Path,
		})
		return
	}

	c.JSON(http.StatusOK, summary)
}
