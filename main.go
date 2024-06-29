package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"urlshortener/controller"
	"urlshortener/database"
	_ "urlshortener/docs"
	"urlshortener/factory"
	"urlshortener/service"
)

func main() {
	route := gin.Default()
	database.ConnectDatabase()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	urlController := controller.URLController{}

	urlVisitRepo := database.URLVisitRepositoryImpl{}
	urlVisitFactory := factory.URLVisitFactory{}
	urlService := service.UrlService{}
	urlVisitService := service.NewURLVisitService(urlService, urlVisitRepo, urlVisitFactory)

	urlVisitController := controller.URLVisitController{
		URLVisitService: urlVisitService,
	}

	//Instantiate URL CRUD methods urls
	route.POST("/urls", urlController.AddURL)
	route.GET("/urls", urlController.GetURLs)
	route.GET("/urls/:externalId/visits", urlVisitController.CountURLVisits)
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Start server in new routine to avoid blocking
	go func() {
		err := route.Run(":8080")
		if err != nil {
			panic(err)
		}

	}()

	// Start new router to serve the redirects
	redirectRouter := gin.Default()
	redirectRouter.GET("/:externalId", urlVisitController.Redirect)
	err := redirectRouter.Run("localhost:8081")
	if err != nil {
		panic(err)
	}
}
