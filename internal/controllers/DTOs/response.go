package dtos

type GetUserTimelineResponse struct {
	Messages []string `json:"messages"`
}

type GetUserResponse struct {
	Name         string   `json:"username"`
	Following    []string `json:"following"`
	Publications []string `json:"publications"`
}
