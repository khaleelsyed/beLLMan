package storage

import (
	"testing"
	"time"

	"github.com/khaleelsyed/beLLMan/internal/types"
)

func setupMockStorage() (*MockStorage, error) {
	storage, err := NewMockStorage()
	if err != nil {
		return nil, err
	}

	if err = storage.Init(); err != nil {
		return nil, err
	}
	return storage, nil
}

func TestMockStoreListChats(t *testing.T) {
	expectedResult := []types.Chat{
		{
			ID:         1,
			Title:      "Python Tutor",
			UpdatedAt:  time.Date(2025, 07, 02, 23, 55, 00, 00, time.UTC),
			MessageIDs: []int{1, 2, 3},
		},
		{ID: 2, Title: "Science Expertise", UpdatedAt: time.Date(2025, 07, 03, 23, 55, 00, 00, time.UTC), MessageIDs: []int{4, 5, 6}},
	}

	storage, err := setupMockStorage()
	if err != nil {
		t.Fatalf("Failed to setup mock storage: %v", err)
	}

	chats, err := storage.ListChats()
	if err != nil {
		t.Fatalf("Failed to list chats: %v", err)
	}

	if len(chats) == 0 {
		t.Fatal("Expected non-empty chat list, got empty")
	}

	for _, chat := range chats {
		found := false
		for _, expected := range expectedResult {
			if chatsEqual(chat, expected) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected chat found: %+v", chat)
		}
	}
}

func TestMockStoreGetChat(t *testing.T) {
	storage, err := setupMockStorage()
	if err != nil {
		t.Fatalf("Failed to setup mock storage: %v", err)
	}

	tests := []struct {
		chatID       int
		expectedChat types.FullChat
		expectError  bool
	}{
		{
			chatID: 1,
			expectedChat: types.FullChat{
				ID:        1,
				Title:     "Python Tutor",
				UpdatedAt: time.Date(2025, 07, 02, 23, 55, 00, 00, time.UTC),
				Messages: []types.Message{
					allMessages[1],
					allMessages[2],
					allMessages[3],
				},
			},
			expectError: false,
		},
		{
			chatID:      999,
			expectError: true,
		},
	}

	for _, test := range tests {
		result, err := storage.GetChat(test.chatID)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error for chat ID %d, got none", test.chatID)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for chat ID %d: %v", test.chatID, err)
			continue
		}

		if !fullChatEqual(result, test.expectedChat) {
			t.Errorf("Expected chat %d to match, but got %+v. Expected %+v", test.chatID, result, test.expectedChat)
		}
	}
}
