package models

// ReceiveMessage struct
type ReceiveMessage struct {
	UpdateID int `json:"update_id"`
}

// ReceiveCiv6PlayByCloud struct
type ReceiveCiv6PlayByCloud struct {
	GameName       string `json:"value1"`
	PlayerName     string `json:"value2"`
	GameTurnNumber string `json:"value3"`
}
