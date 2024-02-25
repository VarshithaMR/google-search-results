package googlesearch

import (
	"fmt"
	"google-search/service/models"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

const (
	mockResponse = "test-util/google-search-response.json"
	mockURL      = "https://mockurl.com"
)

func setUpMockURL() {
	restyClient := resty.New()
	mockresponseFile := mockResponse
	jsonResponse, err := os.ReadFile(mockresponseFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, mockURL, httpmock.NewStringResponder(200, string(jsonResponse)))
}

func expectedResponses() []*models.GoogleSearchResponse {
	response1 := &models.GoogleSearchResponse{
		Kind: "",
		URL: &models.URL{
			Type:     "",
			Template: "",
		},
		SearchInformation: &models.SearchInformation{
			SearchTime:   5.5,
			TotalResults: "444",
		},
		ResponseItems: []*models.ResponseItem{
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
		},
	}
	response2 := &models.GoogleSearchResponse{
		Kind: "",
		URL: &models.URL{
			Type:     "",
			Template: "",
		},
		SearchInformation: &models.SearchInformation{
			SearchTime:   5.5,
			TotalResults: "444",
		},
		ResponseItems: []*models.ResponseItem{
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
			{
				Kind:  "",
				Title: "",
				Link:  "",
			},
		},
	}

	return []*models.GoogleSearchResponse{response1, response2}
}
