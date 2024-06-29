package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"urlshortener/model"
	"urlshortener/service"
	"urlshortener/urlErrors"
)

// swagger embed files

type URLController struct {
	urlService       service.UrlService
	urlVisitsService service.URLVisitService
}

// @BasePath /api/v1

// GetURLs GetUrls godoc
// @Summary get all available short urls
// @Schemes
// @Description get all available short urls
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {array} model.URL
// @Router /urls/ [get]
// TODO add authentication to get urls only for specific user
func (urlController *URLController) GetURLs(context *gin.Context) {
	urls, err := urlController.urlService.GetAll()

	if handleGenericError(context, err) {
		return
	}
	context.JSON(http.StatusOK, urls)
}

// @BasePath /api/v1

// AddURL adds a new URL
// @Summary Add a new URL
// @Description Add a new URL
// @ID add-url
// @Accept  json
// @Produce  json
// @Param   url  body  model.NewURL  true  "URL to be added"
// @Success 200 {object} model.URL
// @Failure 400 {object} urlErrors.GenericError
// @Router /urls [post]
func (urlController *URLController) AddURL(context *gin.Context) {

	var newURL model.NewURL
	err := context.ShouldBindBodyWith(&newURL, binding.JSON)

	if handleAddURLValidationError(context, err) {
		return
	}

	url, err := urlController.urlService.Add(newURL)

	if handleAddURLServiceError(context, err, newURL) {
		return
	}

	context.JSON(http.StatusOK, url)
}

// Error handlers
func handleAddURLServiceError(context *gin.Context, err error, newURL model.NewURL) bool {
	if err != nil {
		log.Printf("Error during adding the URL with message: %v for URL: %v\n", err, newURL)
		//TODO error message is missing
		context.AbortWithStatusJSON(400, urlErrors.GenericError{StatusCode: 400, Err: err})
		return true
	}
	return false
}

func handleAddURLValidationError(context *gin.Context, err error) bool {
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("Error Adding new URL with error: %v\n", err)
		return true

	}
	return false
}

func handleGenericError(context *gin.Context, err error) bool {
	if err != nil {
		log.Printf("Error during getting the URLS %v\n", err)
		context.AbortWithStatusJSON(400, "Bad request")
		return true
	}
	return false
}
