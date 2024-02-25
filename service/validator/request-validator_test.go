package validator

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"google-search/service/models"
)

const (
	urlPath = "/google-results/top-list"
)

func TestValidateRequest(test *testing.T) {
	var size = new(int)
	*size = 10

	handlerRequestWithSize := &models.HandlerRequest{
		Query:          "Black Mirror Bandersnatch",
		ResultQuantity: size,
	}

	handlerRequestWithoutSize := &models.HandlerRequest{
		Query: "How to create Programmable Search Engine?",
	}

	handlerRequestEmpty := &models.HandlerRequest{}

	tests := []struct {
		name             string
		request          *http.Request
		expectedResponse *models.HandlerRequest
		err              error
	}{
		{
			name:             "Get request with user input size",
			request:          createRequest(http.MethodGet, "test-util/request-with-user-input-size.json", urlPath),
			expectedResponse: handlerRequestWithSize,
			err:              nil,
		},

		{
			name:             "Get request with default size",
			request:          createRequest(http.MethodGet, "test-util/request-with-default-size.json", urlPath),
			expectedResponse: handlerRequestWithoutSize,
			err:              nil,
		},

		{
			name:             "Get request with empty body",
			request:          createRequest(http.MethodGet, "test-util/request-empty.json", urlPath),
			expectedResponse: handlerRequestEmpty,
			err:              errors.New("error getting request body"),
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(tt *testing.T) {

			tt.Parallel()
			req, err := ValidateRequest(t.request.Body)
			if err != nil {
				assert.Equal(tt, t.err, err)
			} else {
				assert.Equal(tt, t.expectedResponse, req)
			}
		})
	}
}

func createRequest(httpMethod string, requestJson string, url string) *http.Request {
	reqJson, err := os.ReadFile(requestJson)
	if err != nil {
		return nil
	}

	reader := bytes.NewBuffer(reqJson)
	req, err := http.NewRequest(httpMethod, url, reader)
	if err != nil {
		fmt.Println("Error", err)
		return req
	}

	return req
}
