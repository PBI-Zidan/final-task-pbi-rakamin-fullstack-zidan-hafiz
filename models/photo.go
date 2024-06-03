package models

import "time"

type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(100)" json:"title"`
	Caption   string    `gorm:"type:varchar(255)" json:"caption"`
	PhotoURL  string    `gorm:"type:varchar(255)" json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
