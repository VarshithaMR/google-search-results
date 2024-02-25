package builder

import (
	"reflect"
	"testing"

	"google-search/service/builder/model"
)

func TestBuilderResponse(test *testing.T) {
	response := model.APIResponse{
		Message: "Message",
	}

	tests := []struct {
		name             string
		message          string
		expectedResponse model.APIResponse
	}{
		{
			name:             "Test1 - with any random message",
			message:          "Message",
			expectedResponse: response,
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(tt *testing.T) {
			test.Parallel()
			if actual := BuildResponse(t.message); !reflect.DeepEqual(actual, t.expectedResponse) {
				tt.Errorf("BuildResponse() = %v, expectedResponse %v", actual, t.expectedResponse)
			}
		})
	}
}
