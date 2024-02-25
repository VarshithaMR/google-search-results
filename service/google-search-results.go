package service

import (
	"errors"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"

	"google-search/service/builder"
	"google-search/service/models"
	"google-search/service/providers/googlesearch"
	"google-search/service/validator"
)

const (
	contentType = "Content-Type"
	appType     = "application/json"
)

type SearchCompositionHandler interface {
	GetGoogleSearchResults(http.ResponseWriter, *http.Request, int)
}
type Providers struct {
	GoogleSearchClient googlesearch.GoogleSearchClient
}

func (p *Providers) GetGoogleSearchResults(writer http.ResponseWriter, request *http.Request, resultQuantity int) {
	serviceRequest, err := validator.ValidateRequest(request.Body)
	if err != nil {
		WriteResponse(writer, builder.BuildResponse("❌ Cannot build Proper response"), http.StatusBadRequest)
		return
	}

	response, err := googleSearchResults(serviceRequest, resultQuantity, p.GoogleSearchClient)
	if err != nil {
		WriteResponse(writer, builder.BuildResponse("❌ Cannot build Proper response"), http.StatusInternalServerError)
		return
	}

	WriteResponse(writer, response, http.StatusOK)
	log.Println("💃🏻 ✅ Google Search Process successful")
}

func googleSearchResults(serviceRequest *models.HandlerRequest, resultQuantity int, provider googlesearch.GoogleSearchClient) (*models.HandlerResponse, error) {
	var (
		quantity      int
		finalResponse *models.HandlerResponse
	)

	if serviceRequest.ResultQuantity == nil {
		quantity = resultQuantity
	} else {
		quantity = *serviceRequest.ResultQuantity
	}

	results := provider.GetSearchResults(serviceRequest.Query)
	if len(results.ResponseItems) < quantity {
		log.Println("😣 Less no of results than expected")
		finalResponse = mapGResultsToService(results, len(results.ResponseItems))
		return finalResponse, nil
	}

	if len(results.ResponseItems) >= quantity {
		finalResponse = mapGResultsToService(results, quantity)
		return finalResponse, nil
	}

	return nil, errors.New("❌ not able to form response")
}

func mapGResultsToService(googleResults *models.GoogleSearchResponse, quantity int) *models.HandlerResponse {
	var (
		responseItem  *models.HandlerResponseItem
		responseItems []*models.HandlerResponseItem
	)

	for _, response := range googleResults.ResponseItems[:quantity] {
		responseItem.Title = response.Title
		responseItem.Link = response.Link
		responseItems = append(responseItems, responseItem)
	}

	return &models.HandlerResponse{
		ResponseTime:  googleResults.SearchInformation.SearchTime,
		ResponseItems: responseItems,
	}
}
func WriteResponse(rw http.ResponseWriter, resp interface{}, responseCode int) {
	rw.WriteHeader(responseCode)
	rw.Header().Set(contentType, appType)
	bytes, _ := jsoniter.Marshal(resp)
	rw.Write(bytes)
}

func NewSearchCompositionHandler(providers Providers) SearchCompositionHandler {

	return &Providers{
		GoogleSearchClient: providers.GoogleSearchClient,
	}
}
