package model

type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"` // hashed
	CreatedAt int64  `json:"created_at"`
}
