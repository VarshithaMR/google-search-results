package googlesearch

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	mockResponse10 = "test-util/google-search-10.json"
	mockResponse5  = "test-util/google-search-5.json"
	mockQuery      = "mockQuery"
)

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

func TestSearch_GetSearchResults_10(t *testing.T) {
	restyClient := resty.New()
	jsonResponse, err := os.ReadFile(mockResponse10)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	recorder := httptest.NewRecorder()
	reader := bytes.NewReader(jsonResponse)
	_, err = recorder.Body.ReadFrom(reader)
	if err != nil {
		t.Fatal(err)
	}
	resp := recorder.Result()

	httpmock.RegisterResponder(http.MethodGet, mockURLEndpoint, httpmock.ResponderFromResponse(resp))

	client := GoogleMockClient(restyClient)
	result, err := client.GetGoogleSearchResults(mockQuery, 10)
	fmt.Println(result)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.ResponseItems)
	assert.Equal(t, 10, len(result.ResponseItems))
}

func TestSearch_GetSearchResults_5(t *testing.T) {
	restyClient := resty.New()
	jsonResponse, err := os.ReadFile(mockResponse5)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	recorder := httptest.NewRecorder()
	reader := bytes.NewReader(jsonResponse)
	_, err = recorder.Body.ReadFrom(reader)
	if err != nil {
		t.Fatal(err)
	}
	resp := recorder.Result()

	httpmock.RegisterResponder(http.MethodGet, mockURLEndpoint, httpmock.ResponderFromResponse(resp))

	client := GoogleMockClient(restyClient)
	result, err := client.GetGoogleSearchResults(mockQuery, 5)
	fmt.Println(result)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.ResponseItems)
	assert.Equal(t, 5, len(result.ResponseItems))
}

func TestSearch_GetSearchResults_Empty(t *testing.T) {
	restyClient := resty.New()
	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(http.MethodGet, mockURLEndpoint, nil)

	client := GoogleMockClient(restyClient)
	result, err := client.GetGoogleSearchResults(mockQuery, 5)
	fmt.Println(result)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
}
