package model

import "time"

// User struct to represent user details
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Score     int       `json:"score"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
