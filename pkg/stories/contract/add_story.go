package contract

type AddStoryRequest struct {
	Story
}

type AddStoryResponse struct {
	StoryID string `json:"story_id"`
}
