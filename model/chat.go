package model

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatSender struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ChatResponse struct {
	Sender  ChatSender `json:"sender"`
	Message string     `json:"message"`
}
