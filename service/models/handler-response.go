package models

type HandlerResponse struct {
	ResponseTime  float64                `json:"queryTime"`
	ResponseItems []*HandlerResponseItem `json:"top links"`
}

type HandlerResponseItem struct {
	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`
}
