package contract

type DeleteStoryRequest struct {
	ID string `json:"id"`
}

type DeleteStoryResponse struct {
	Success bool `json:"success"`
}
