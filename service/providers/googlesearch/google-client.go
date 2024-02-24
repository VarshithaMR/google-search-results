package googlesearch

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"google-search/service/models"
)

const (
	contentType   = "Content-Type"
	appType       = "application/json"
	endpoint      = "/customsearch/v1"
	paramQuery    = "q"
	paramKey      = "key"
	paramCustomId = "cx"

	envVarGoogleSearchURL      = "GOOGLE_SEARCH_URL"
	envVarGoogleAPIKey         = "GOOGLE_API_KEY"
	envVarCustomSearchEngineId = "GOOGLE_CUSTOM_SEARCH_ID"
)

type GoogleSearchClient interface {
	GetSearchResults(string) *models.GoogleSearchResponse
}

type search struct {
	httpClient *resty.Client
}

func (g *search) GetSearchResults(query string) *models.GoogleSearchResponse {
	baseUrl := g.httpClient.BaseURL
	response, err := g.httpClient.R().
		SetHeader(contentType, appType).
		SetQueryParam(paramQuery, query).
		SetResult(models.GoogleSearchResponse{}).
		Get(baseUrl + endpoint)
	if err != nil {
		log.Warnf("❌ Google API Search error: %s", err)
	}

	return response.Result().(*models.GoogleSearchResponse)
}

func NewGoogleClient(properties *viper.Viper) GoogleSearchClient {
	url := properties.GetString(envVarGoogleSearchURL)
	apiKey := properties.GetString(envVarGoogleAPIKey)
	searchEngineId := properties.GetString(envVarCustomSearchEngineId)

	searchClient := resty.New()

	searchClient.SetBaseURL(url)
	searchClient.SetQueryParam(paramKey, apiKey)
	searchClient.SetQueryParam(paramCustomId, searchEngineId)

	return &search{
		httpClient: searchClient,
	}
}
