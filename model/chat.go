package model

type ChatIncoming struct {
	Message string `json:"message"`
}

type ChatSender struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ChatOutgoing struct {
	Sender  ChatSender `json:"sender"`
	Message string     `json:"message"`
}
