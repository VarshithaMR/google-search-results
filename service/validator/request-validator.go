package validator

import (
	"encoding/json"
	"io"

	log "github.com/sirupsen/logrus"

	"google-search/service/models"
)

func ValidateRequest(body io.ReadCloser) (requestBody *models.HandlerRequest, err error) {
	decoder := json.NewDecoder(body)

	if err := decoder.Decode(&requestBody); err != nil {
		log.Error("Error getting request body %s", err)
		return nil, err
	}
	return requestBody, nil
}
