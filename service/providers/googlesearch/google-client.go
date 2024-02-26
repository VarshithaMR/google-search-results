package googlesearch

import (
	"encoding/json"
	"fmt"
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

	paramNum = "num"

	envVarGoogleSearchURL      = "GOOGLE_SEARCH_URL"
	envVarGoogleAPIKey         = "GOOGLE_API_KEY"
	envVarCustomSearchEngineId = "GOOGLE_CUSTOM_SEARCH_ID"
)

type GoogleSearchClient interface {
	GetSearchResults(string, int) (*models.GoogleSearchResponse, error)
}

type search struct {
	httpClient *resty.Client
}

func (g *search) GetSearchResults(query string, quantity int) (*models.GoogleSearchResponse, error) {
	baseUrl := g.httpClient.BaseURL
	response, err := g.httpClient.R().
		SetHeader(contentType, appType).
		SetQueryParam(paramQuery, query).
		SetQueryParam(paramNum, string(rune(quantity))).
		SetResult(models.GoogleSearchResponse{}).
		Get(baseUrl + endpoint)
	if err != nil {
		log.Warnf("❌ Google API Search error: %s", err)
		return nil, err
	}

	// for mocking- test cases
	if baseUrl == mockURL {
		var x models.GoogleSearchResponse
		err = json.Unmarshal(response.Body(), &x)
		if err != nil {
			fmt.Println("Testcase : Unmarshal err", err)
		}
		return &x, nil
	}

	return response.Result().(*models.GoogleSearchResponse), nil
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
