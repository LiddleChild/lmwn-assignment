package main

import (
	"fmt"
	"log"

	"github.com/LiddleChild/covid-stat/config"
	"github.com/LiddleChild/covid-stat/internal/summary"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()

	if config.IsDevelopment() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	summaryRepo := summary.NewRepository()
	summaryService := summary.NewService(summaryRepo, config)
	summaryHandler := summary.NewHandler(summaryService)

	r.GET("/covid/summary", summaryHandler.GetSummary)

	// prevent macos from asking permission for accepting incoming network connections
	host := ""
	if config.IsDevelopment() {
		host = "localhost"
	}

	err := r.Run(fmt.Sprintf("%v:%v", host, config.AppPort))
	if err != nil {
		log.Fatal(err)
	}
}
