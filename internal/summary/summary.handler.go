package summary

import (
	"net/http"

	"github.com/LiddleChild/covid-stat/internal/covid_case"
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
	var cases []covid_case.CovidCase

	err := h.service.GetSummary(&cases)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, cases)
}
