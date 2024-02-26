package googlesearch

import (
	"github.com/go-resty/resty/v2"
)

const (
	mockURL         = "https://mockurl.com"
	mockURLEndpoint = "https://mockurl.com/customsearch/v1"
	mockKey         = "mockapikey"
	mockId          = "mockengineid"
)

func GoogleMockClient(searchClient *resty.Client) GoogleSearchClient {
	searchClient.SetBaseURL(mockURL)
	searchClient.SetQueryParam(paramKey, mockKey)
	searchClient.SetQueryParam(paramCustomId, mockId)
	return &search{
		httpClient: searchClient,
	}
}
