package service

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"google-search/service/models"
	"google-search/service/providers/googlesearch"
)

func TestNewSearchCompositionHandler(test *testing.T) {
	var quantity *int
	restyClient := resty.New()
	mockGoogleClient := googlesearch.GoogleMockClient(restyClient)
	jsonResponse, err := os.ReadFile("./service/providers/googlesearch/test-util/google-search-test.json")
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
		test.Fatal(err)
	}
	resp := recorder.Result()

	httpmock.RegisterResponder(http.MethodGet, "https://mockurl.com/customsearch/v1", httpmock.ResponderFromResponse(resp))

	*quantity = 2
	request1 := &models.HandlerRequest{
		Query:          "First query",
		ResultQuantity: quantity,
	}

	response1 := &models.HandlerResponse{
		ResponseTime: 1.5,
		ResponseItems: []*models.HandlerResponseItem{
			{
				Title: "First query",
				Link:  "Link 1",
			},

			{
				Title: "First query",
				Link:  "Link 2",
			},
		},
	}

	tests := []struct {
		name             string
		request          *models.HandlerRequest
		defaultQuantity  int
		expectedResponse *models.HandlerResponse
		err              error
	}{
		{
			name:             "With user input quantity",
			defaultQuantity:  2,
			request:          request1,
			expectedResponse: response1,
			err:              nil,
		},
		{
			name:             "with nil request body",
			defaultQuantity:  2,
			request:          &models.HandlerRequest{},
			expectedResponse: nil,
			err:              errors.New("❌ empty request"),
		},
		{
			name:             "with user input more than default quantity",
			defaultQuantity:  3,
			request:          &models.HandlerRequest{},
			expectedResponse: response1,
			err:              nil,
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(tt *testing.T) {
			response, err := GoogleSearchResults(t.request, t.defaultQuantity, mockGoogleClient)
			if err != nil {
				assert.Equal(tt, t.err, err)
			} else {
				assert.Equal(tt, t.expectedResponse, response)
			}
		})
	}

}
