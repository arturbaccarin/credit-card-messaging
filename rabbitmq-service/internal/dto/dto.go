package dto

type MessagePayload struct {
	Message string `json:"message"`
}

type MessagesResponse struct {
	Messages []string `json:"messages"`
}
