package models

type HandlerRequest struct {
	Query          string `json:"userQuery"`
	ResultQuantity *int   `json:"resultQuantity,omitempty"`
}
