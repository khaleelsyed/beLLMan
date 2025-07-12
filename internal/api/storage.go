package api

import "github.com/khaleelsyed/beLLMan/internal/types"

type Storage interface {
	Init() error
	ListChats() ([]types.Chat, error)
	GetChat(chatID int) (types.FullChat, error)
}
