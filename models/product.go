package models

import "time"

type Product struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at"`
	Name         string    `json:"name"`
	Expiry       int       `json:"expiry"`
	SerialNumber string    `json:"serial_number"`
}
