package googlesearch

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"

	"google-search/service/models"
)

const (
	mockKey = "mockapikey"
	mockId  = "mockengineid"
	//mockURL = "https://mockurl.com"
)

// TODO: to fix the mocking properly
func TestSearch_GetSearchResults(test *testing.T) {
	setUpMockURL()

	properties := viper.New()
	properties.Set(envVarGoogleSearchURL, mockURL)
	properties.Set(envVarGoogleAPIKey, mockKey)
	properties.Set(envVarCustomSearchEngineId, mockId)

	// Create the GoogleSearchClient instance
	client := NewGoogleClient(properties)

	responses := expectedResponses()
	tests := []struct {
		name             string
		query            string
		size             int
		expectedResponse *models.GoogleSearchResponse
		err              error
	}{
		{
			name:             "Google client responding with 10 items",
			query:            "how to create fb account",
			size:             10,
			expectedResponse: responses[0],
			err:              nil,
		},
		{
			name:             "Google client responding with 5",
			query:            "how to create fb account",
			size:             5,
			expectedResponse: responses[1],
			err:              nil,
		},
		{
			name:             "Google client responding error",
			query:            "how to create fb account",
			size:             3,
			expectedResponse: nil,
			err:              errors.New("error getting request body"),
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := client.GetSearchResults("mockQuery", 10)
			if err != nil {
				assert.Equal(tt, err, t.err)
			} else {
				assert.Equal(tt, result, t.expectedResponse)
			}
		})
	}
}

func TestNewGoogleClient(test *testing.T) {
	properties := viper.New()
	properties.Set(envVarGoogleSearchURL, mockURL) // Use the same mock URL here
	properties.Set(envVarGoogleAPIKey, mockKey)
	properties.Set(envVarCustomSearchEngineId, mockId)

	emptyProperties := viper.New()

	tests := []struct {
		name           string
		url            string
		key            string
		id             string
		expectedClient GoogleSearchClient
	}{
		{
			name: "all values set",
			url:  mockURL,
			key:  mockKey,
			id:   mockId,
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(tt *testing.T) {
			tt.Parallel()
			client := NewGoogleClient(properties)
			assert.Equal(tt, t.url, client.(*search).httpClient.BaseURL)
			assert.Equal(tt, t.key, client.(*search).httpClient.QueryParam.Get(paramKey))
			assert.Equal(tt, t.id, client.(*search).httpClient.QueryParam.Get(paramCustomId))

			emptyClient := NewGoogleClient(emptyProperties)
			assert.Empty(tt, emptyClient.(*search).httpClient.BaseURL)
			assert.Empty(tt, emptyClient.(*search).httpClient.QueryParam.Get(paramKey))
			assert.Empty(tt, emptyClient.(*search).httpClient.QueryParam.Get(paramCustomId))
		})
	}
}

/*func expectedResponses() []*models.GoogleSearchResponse {
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
}*/
