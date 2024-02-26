package writer

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"google-search/service/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteResponse(test *testing.T) {
	response := &models.HandlerResponse{
		ResponseTime: 5.45,
		ResponseItems: []*models.HandlerResponseItem{
			{
				Title: "Search_Query",
				Link:  "Link1",
			},
		},
	}

	tests := []struct {
		name             string
		expectedResponse *models.HandlerResponse
		statusCode       int
		contentType      string
	}{
		{
			name:             "Success response",
			expectedResponse: response,
			statusCode:       http.StatusOK,
			contentType:      appType,
		},

		{
			name:             "Success response",
			expectedResponse: response,
			statusCode:       http.StatusBadRequest,
			contentType:      appType,
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(tt *testing.T) {
			tt.Parallel()
			recorder := httptest.NewRecorder()                        // mock http.Responsewriter
			WriteResponse(recorder, t.expectedResponse, t.statusCode) //call response writer

			assert.Equal(tt, t.statusCode, recorder.Code)
			assert.Equal(tt, appType, recorder.Header().Get(contentType))

			_ = json.Unmarshal(recorder.Body.Bytes(), t.expectedResponse)
			assert.Equal(test, t.expectedResponse, response)
		})
	}

}
