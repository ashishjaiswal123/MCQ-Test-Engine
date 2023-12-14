package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User model to represent user details and scores
type User struct {
	gorm.Model
	Name      string
	Email     string
	Score     int
	StartTime time.Time
	EndTime   time.Time
}
