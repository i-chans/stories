package contract

type SearchStoryRequest struct {
	Query string `json:"query"`
}

type SearchStoryResponse struct {
	Stories []Story `json:"stories"`
}
