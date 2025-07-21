package dtos

type MessageResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type GetUserTimelineResponse struct {
	Feed []MessageResponse `json:"feed"`
}

type GetUserResponse struct {
	Name         string   `json:"username"`
	Following    []string `json:"following"`
	Publications []string `json:"publications"`
}
