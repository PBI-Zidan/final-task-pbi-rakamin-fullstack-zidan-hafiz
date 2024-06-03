package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(100);unique_index" json:"username"`
	Email     string    `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string    `json:"password"`
	Photo     Photo     `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL;" json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
