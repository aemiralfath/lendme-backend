package models

import (
	"time"
)

type Role struct {
	RoleID    int       `json:"role_id" db:"role_id" binding:"omitempty"`
	Name      string    `json:"name" db:"name" binding:"omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
