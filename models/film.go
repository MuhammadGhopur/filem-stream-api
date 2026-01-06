package models

import "time"

type Film struct {
	ID uint `gorm:"primaryKey"`

	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	VideoURL    string `gorm:"column:video_url"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
