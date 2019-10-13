package models

import (
	"time"
)

// Product represent the product model
type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Sku       string    `json:"sku"`
	Type      string    `json:"type"`
	Category  int64     `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
