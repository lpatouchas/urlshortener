package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	//1. load .env vars
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//2. initialize db
	database.ConnectDatabase()

	//3. set log flags
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//4. initialize repos
	urlVisitRepo := &database.URLVisitRepositoryImpl{}
	urlVisitFactory := factory.URLVisitFactory{}
	urlRepository := &database.URLRepository{}

	//5. Initialize services
	urlService := service.NewUrlService(urlRepository, urlVisitRepo)
	urlVisitService := service.NewURLVisitService(&urlService, urlVisitRepo, urlVisitFactory)

	//6. Initialize controllers
	urlController := controller.NewURLController(urlService, urlVisitService)
	urlVisitController := controller.URLVisitController{
		URLVisitService: urlVisitService,
	}

	//7. Instantiate and start URL CRUD methods urls
	route := gin.Default()
	route.POST("/urls", urlController.AddURL)
	route.GET("/urls", urlController.GetURLs)
	route.GET("/urls/:externalId/visits", urlVisitController.CountURLVisits)
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//8. start server in new routine to avoid blocking
	go func() {
		err := route.Run(":8080")
		if err != nil {
			panic(err)
		}

	}()

	//9. Start new router to serve the redirects
	redirectRouter := gin.Default()
	redirectRouter.GET("/:externalId", urlVisitController.Redirect)
	err = redirectRouter.Run("localhost:8081")
	if err != nil {
		panic(err)
	}
}
