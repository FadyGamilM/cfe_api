package domain

import "time"

type Coffee struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Region     string    `json:"region"`
	RoastLevel string    `json:"roast_level"`
	Price      float64   `json:"price"`
	GrindUnit  string    `json:"grind_unit"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
