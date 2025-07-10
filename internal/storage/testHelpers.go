package storage

import (
	"slices"

	"github.com/khaleelsyed/beLLMan/internal/types"
)

func chatsEqual(a, b types.Chat) bool {
	return a.ID == b.ID &&
		a.Title == b.Title &&
		a.UpdatedAt.Equal(b.UpdatedAt) &&
		slices.Equal(a.MessageIDs, b.MessageIDs)
}

func messagesEqual(a, b types.FullChat) bool {
	return a.ID == b.ID &&
		a.Title == b.Title &&
		a.UpdatedAt.Equal(b.UpdatedAt) &&
		slices.Equal(a.Messages, b.Messages)
}
