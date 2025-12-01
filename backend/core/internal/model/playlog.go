package model

type PlayLog struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	BoardGameID int64  `json:"board_game_id"`
	PlayedAt    int64  `json:"played_at"`
	Duration    int    `json:"duration"` // in minutes
	Winner      string `json:"winner"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}
