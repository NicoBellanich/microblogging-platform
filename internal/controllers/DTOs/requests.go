package dtos

type FollowRequest struct {
	UserName  string `json:"user_id"`
	NewFollow string `json:"new_follow"`
}

type GetUserTimelineRequest struct {
	UserName string `json:"user_id"`
}

type CreateUserRequest struct {
	UserName string `json:"user_id"`
}

type GetUserRequest struct {
	UserName string `json:"user_id"`
}

type CreatePublicationRequest struct {
	UserName string `json:"user_id"`
	Content  string `json:"content"`
}
