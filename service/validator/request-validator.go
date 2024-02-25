package validator

import (
	"encoding/json"
	"errors"
	"io"

	"google-search/service/models"
)

func ValidateRequest(body io.ReadCloser) (requestBody *models.HandlerRequest, err error) {
	decoder := json.NewDecoder(body)

	if err := decoder.Decode(&requestBody); err != nil {
		//log.Error("Error getting request body:", err)
		return nil, errors.New("error getting request body")
	}
	return requestBody, nil
}
