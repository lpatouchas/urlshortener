package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"urlshortener/service"
	"urlshortener/urlErrors"
)

type URLVisitController struct {
	URLVisitService service.URLVisitService //TODO fully understand why need pointer
}

func (urlVisitController *URLVisitController) Redirect(context *gin.Context) {
	externalId := context.Param("externalId")

	if externalId == "" {
		context.IndentedJSON(http.StatusBadRequest, "{'status':'missing externalid'}")
	}

	redirectorURL, err := urlVisitController.URLVisitService.GetRedirectorURL(externalId)

	if handlerRedirectError(context, err, externalId) {
		return
	}

	context.Redirect(http.StatusFound, redirectorURL)
}

// CountURLVisits counts the number of visits for a URL
// @Summary Count URL visits
// @Description Count the number of visits for a URL based on its external ID
// @ID count-url-visits
// @Produce json
// @Param externalId path string true "External ID of the URL"
// @Success 200 {model.URLVisits} map[string]int "visits"
// @Failure 400 {object} urlErrors.GenericError
// @Failure 404 {object} urlErrors.GenericError
// @Router /urls/{externalId}/visits [get]
func (urlVisitController *URLVisitController) CountURLVisits(context *gin.Context) {
	externalId := context.Param("externalId")

	visits, err := urlVisitController.URLVisitService.CountURLVisits(externalId)
	handleGenericError(context, err)
	context.JSON(http.StatusOK, visits)
}

// Error handlers
func handlerRedirectError(context *gin.Context, err error, externalId string) bool {
	if err != nil {
		log.Printf("Error during getting the Redirector URL for external id %s %v\n", externalId, err)

		//TODO need to further understand the pointer here
		var urlErr *urlErrors.RedirectError
		if errors.As(err, &urlErr) {
			fmt.Printf("An error occured: %s\n", urlErr.UrlExternalId)
			context.AbortWithStatusJSON(urlErr.StatusCode, urlErr)
		} else {
			context.AbortWithStatusJSON(400, urlErrors.FromExternalID(externalId))
		}

		return true
	}
	return false
}
