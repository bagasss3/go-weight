package main

import (
	//"net/http"

	"github.com/bagasss3/go-weight/internal/controller"
	"github.com/bagasss3/go-weight/internal/db"
	"github.com/bagasss3/go-weight/internal/httpsvc"
	"github.com/bagasss3/go-weight/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	log "github.com/sirupsen/logrus"
)

func main() {
	db.InitializeDbConn()
	sqlDB, err := db.PostgreDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	err = db.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	weightRepo := repository.NewWeighRepository(db.PostgreDB)
	weightController := controller.NewWeightController(weightRepo)

	httpServer := echo.New()
	httpServer.Use(middleware.CORS())
	apiGroup := httpServer.Group("/api")
	httpsvc.RouteService(apiGroup, weightController, weightRepo)
	httpServer.Logger.Fatal(httpServer.Start(":8000"))
}
