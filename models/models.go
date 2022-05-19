package models

//ResponseObject to construct response
type ResponseObject struct {
	Success bool `json:"success,omitempty"`
}

type WebCrawlResponse struct {
	WebCrawlResult []WebCrawlResult `json:"result,omitempty"`
}

type WebCrawlResult struct {
	WebUrl string `json:"url,omitempty"`
	Data   string `json:"data,omitempty"`
}

type IncomingRequest struct {
	WebUrl []string `json:"urls"`
}
