package main

import (
	"log"

	"github.com/LiddleChild/covid-stat/internal/summary"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	summaryRepo := summary.NewRepository()
	summaryService := summary.NewService(summaryRepo)
	summaryHandler := summary.NewHandler(summaryService)

	r.GET("/covid/summary", summaryHandler.GetSummary)

	err := r.Run("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
}
