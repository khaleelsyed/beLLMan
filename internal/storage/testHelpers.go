package storage

import (
	"github.com/khaleelsyed/beLLMan/internal/types"
)

func chatsEqual(a, b types.Chat) bool {
	return a.ID == b.ID &&
		a.Title == b.Title &&
		a.UpdatedAt.Equal(b.UpdatedAt)
}

func allMessagesEqual(a, b []types.Message) bool {
	for _, message := range a {
		found := false
		for _, compareMessage := range b {
			if messageEqual(message, compareMessage) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func messageEqual(a, b types.Message) bool {
	return a.ID == b.ID &&
		a.Role == b.Role &&
		a.Content == b.Content &&
		a.CreatedAt.Equal(b.CreatedAt)
}
