package service

type AddUserRequest struct {
	UserID string         `json:"user_id" swaggerignore:"true"`
	Input  AddUserPayload `json:"input"`
}

type AddUserPayload struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}
