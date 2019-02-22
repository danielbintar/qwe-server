package model

type MoveIncoming struct {
	Direction string `json:"direction"`
}

type MoveCharacter struct {
	ID uint `json:"id"`
}

type MoveOutgoing struct {
	Character MoveCharacter `json:"character"`
	X         uint          `json:"x"`
	Y         uint          `json:"y"`
}
