package service

import (
	jsoniter "github.com/json-iterator/go"
	"google-search/service/builder"
	"google-search/service/builder/model"
	"net/http"

	"google-search/service/providers/googlesearch"
)

const (
	contentType = "Content-Type"
	appType     = "application/json"
)

type SearchCompositionHandler interface {
	GetGoogleSearchResults(http.ResponseWriter, *http.Request, int)
}
type Providers struct {
	GoogleSearchClient googlesearch.GoogleSearchInterface
}

func (p *Providers) GetGoogleSearchResults(writer http.ResponseWriter, request *http.Request, resultQuantity int) {

	WriteResponse(writer, builder.BuildResponse("Google Search Process successful"), http.StatusOK)
}

func WriteResponse(rw http.ResponseWriter, resp interface{}, responseCode int) {
	rw.WriteHeader(responseCode)
	rw.Header().Set(contentType, appType)
	bytes, _ := jsoniter.Marshal(resp)
	rw.Write(bytes)
}

func BuildResponse(message string) model.APIResponse {
	response := model.APIResponse{
		Message: message,
	}
	return response
}
