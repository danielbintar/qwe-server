package model

type WebsocketRequest struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}
