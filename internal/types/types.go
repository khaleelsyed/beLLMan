package types

import "time"

type Chat struct {
	ID         int       `json:"id" validate:"required"`
	Title      string    `json:"title" validate:"required"`
	UpdatedAt  time.Time `json:"updated_at" validate:"required"`
	MessageIDs []int     `json:"message_ids" validate:"required"`
}

type Message struct {
	ID      int       `json:"id" validate:"required"`
	Role    Role      `json:"role" validate:"required"`
	Content string    `json:"content" validate:"required"`
	Sent    time.Time `json:"sent" validate:"required"`
}

type FullChat struct {
	ID        int       `json:"id" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
	Messages  []Message `json:"messages" validate:"required"`
}
