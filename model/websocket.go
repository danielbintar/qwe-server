package model

import "encoding/json"

type WebsocketRequest struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}
