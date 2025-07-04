package types

import "fmt"

type ErrChatNotFound struct {
	ChatID int
}

func (e *ErrChatNotFound) Error() string {
	return fmt.Sprintf("chat ID not found: %d", e.ChatID)
}
