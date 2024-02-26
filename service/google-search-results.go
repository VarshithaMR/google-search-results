package service

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"

	"google-search/service/builder"
	"google-search/service/models"
	"google-search/service/providers/googlesearch"
	"google-search/service/validator"
	"google-search/service/writer"
)

type SearchCompositionHandler interface {
	GetSearchResults(http.ResponseWriter, *http.Request, int)
}
type Providers struct {
	GoogleSearchClient googlesearch.GoogleSearchClient
}

func (p *Providers) GetSearchResults(rw http.ResponseWriter, request *http.Request, resultQuantity int) {
	serviceRequest, err := validator.ValidateRequest(request.Body)
	if err != nil {
		writer.WriteResponse(rw, builder.BuildResponse("❌ Cannot build Proper response"), http.StatusBadRequest)
		return
	}

	response, err := GoogleSearchResults(serviceRequest, resultQuantity, p.GoogleSearchClient)
	if err != nil {
		writer.WriteResponse(rw, builder.BuildResponse("❌ Cannot build Proper response"), http.StatusInternalServerError)
		return
	}

	writer.WriteResponse(rw, response, http.StatusOK)
	log.Println("💃🏻 ✅ Google Search Process successful")
}

func GoogleSearchResults(serviceRequest *models.HandlerRequest, resultQuantity int, provider googlesearch.GoogleSearchClient) (*models.HandlerResponse, error) {
	var (
		quantity      int
		finalResponse *models.HandlerResponse
	)

	if serviceRequest == nil {
		resultError := "❌ empty request"
		return nil, errors.New(resultError)
	}

	if serviceRequest.ResultQuantity == nil {
		//default quantity of items
		quantity = resultQuantity
	} else {
		quantity = *serviceRequest.ResultQuantity
	}

	if quantity > resultQuantity {
		log.Infof("😣 Currently range is valid between 1 to 10")
	}

	//call google api
	results, err := provider.GetGoogleSearchResults(serviceRequest.Query, quantity)
	if err != nil {
		return nil, err
	}

	if results == nil || results.ResponseItems == nil {
		resultError := "❌ not able to get response"
		log.Warnf(resultError)
		return nil, errors.New(resultError)
	}

	finalResponse = mapGResultsToService(results, quantity)
	return finalResponse, nil
}

func mapGResultsToService(googleResults *models.GoogleSearchResponse, quantity int) *models.HandlerResponse {
	var (
		responseItem  *models.HandlerResponseItem
		responseItems []*models.HandlerResponseItem
	)

	for _, response := range googleResults.ResponseItems[:quantity] {
		responseItem = &models.HandlerResponseItem{
			Title: response.Title,
			Link:  response.Link,
		}
		responseItems = append(responseItems, responseItem)
	}

	return &models.HandlerResponse{
		ResponseTime:  googleResults.SearchInformation.SearchTime,
		ResponseItems: responseItems,
	}
}

func NewSearchCompositionHandler(providers Providers) SearchCompositionHandler {

	return &Providers{
		GoogleSearchClient: providers.GoogleSearchClient,
	}
}
