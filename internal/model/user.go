package model

import "time"

// User struct to represent user details
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Score     int       `json:"score"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
