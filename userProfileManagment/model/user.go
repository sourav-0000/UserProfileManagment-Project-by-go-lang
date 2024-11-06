package model

// User represents a user profile
type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name" gorm:"autoIncrement"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
	Address string `json:"address"`
}
