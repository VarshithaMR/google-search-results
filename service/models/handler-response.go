package models

type HandlerResponse struct {
	Query         string   `json:"user query"`
	ResponseTime  string   `json:"queryTime"`
	ResponseItems []string `json:"top links"`
}
