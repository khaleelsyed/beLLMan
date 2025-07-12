package types

import "time"

type Chat struct {
	ID        int       `json:"id" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}

type Message struct {
	ID        int       `json:"id" validate:"required"`
	Role      Role      `json:"role" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}
