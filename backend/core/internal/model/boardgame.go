package model

type Boardgame struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MinPlayers  int     `json:"min_players"`
	MaxPlayers  int     `json:"max_players"`
	PlayTime    int     `json:"play_time"` // in minutes
	Weight      float64 `json:"weight"`    // complexity rating
	Rating      float64 `json:"rating"`
}
