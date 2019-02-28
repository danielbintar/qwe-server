package model

import "encoding/json"

type WebsocketRequest struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type OutgoingMessage struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type LeaveTownData struct {
	ID uint `json:"id"`
}

type LeaveRegionData struct {
	ID uint `json:"id"`
}

type InitBattleData struct {
	ID        uint `json:"id"`
	MonsterID uint `json:"monster_id"`
}
