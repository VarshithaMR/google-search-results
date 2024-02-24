package builder

import "google-search/service/builder/model"

// BuildResponse to build the responses by mapping the message within
func BuildResponse(message string) model.APIResponse {
	response := model.APIResponse{
		Message: message,
	}
	return response
}
