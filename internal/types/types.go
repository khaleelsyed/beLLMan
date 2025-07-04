package types

import "time"

type Chat struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	Messages  []Message `json:"messages"`
}

type Message struct {
	Role    Role      `json:"role"`
	Content string    `json:"content"`
	Sent    time.Time `json:"sent"`
}
