package usecase_dto

type Class struct {
	ID           int    `json:"id"` // primary key
	Classname    string `json:"className"`
	Info         string `json:"info"`
	Announcement string `json:"announcement"`
	RoomCode     string `json:"roomCode"`
	Level        string `json:"level"`
}
