package models

type GoogleSearchResponse struct {
	Kind              string             `json:"kind,omitempty"`
	URL               *URL               `json:"url,omitempty"`
	SearchInformation *SearchInformation `json:"searchInformation,omitempty"`
	ResponseItems     []*ResponseItem    `json:"items,omitempty"`
}

type URL struct {
	Type     string `json:"type,omitempty"`
	Template string `json:"template,omitempty"`
}

type SearchInformation struct {
	SearchTime   float64 `json:"searchTime"`
	TotalResults string  `json:"totalResults,omitempty"`
}

type ResponseItem struct {
	Kind  string `json:"kind,omitempty"`
	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`
}
