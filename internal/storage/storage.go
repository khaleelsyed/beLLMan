package storage

import "github.com/khaleelsyed/beLLMan/internal/types"

type Storage interface {
	ListChats() ([]types.Chat, error)
	GetChat(chatID int) (types.Chat, error)
}
