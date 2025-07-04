package storage

import (
	"time"

	"github.com/khaleelsyed/beLLMan/internal/types"
)

type MockStorage struct{}

var allChats map[int]types.Chat = map[int]types.Chat{
	1: {
		ID:         1,
		Title:      "Python Tutor",
		UpdatedAt:  time.Date(2025, 07, 02, 23, 55, 00, 00, time.UTC),
		MessageIDs: []int{1, 2, 3},
	},
	2: {
		ID:         2,
		Title:      "Science Expertise",
		UpdatedAt:  time.Date(2025, 07, 03, 23, 55, 00, 00, time.UTC),
		MessageIDs: []int{4, 5, 6},
	},
}

var allMessages map[int]types.Message = map[int]types.Message{
	1: {
		ID:      1,
		Role:    types.RoleSystem,
		Content: "You are a Python Tutor AI, dedicated to helping users learn Python and build end-to-end projects using Python and its related libraries. Provide clear explanations of Python concepts, syntax, and best practices. Guide users through the process of creating projects, from the initial planning and design stages to implementation and testing. Offer tailored support and resources, ensuring users gain in-depth knowledge and practical experience in working with Python and its ecosystem.",
		Sent:    time.Date(2025, 07, 02, 23, 50, 00, 00, time.UTC),
	},
	2: {
		ID:      2,
		Role:    types.RoleUser,
		Content: "What is a `SyntaxError`?",
		Sent:    time.Date(2025, 07, 02, 23, 53, 00, 00, time.UTC),
	},
	3: {
		ID:      3,
		Role:    types.RoleAssistant,
		Content: "A **SyntaxError** in Python is an error that occurs when the code you write does not follow the language’s syntax rules, making it impossible for the Python interpreter to understand or execute your program[3][5][6]. These errors are detected during the **parsing stage**, before any code is actually run, and they immediately halt execution until the issue is fixed[1][2][5].",
		Sent:    time.Date(2025, 07, 03, 23, 55, 00, 00, time.UTC),
	},
	4: {
		ID:      4,
		Role:    types.RoleSystem,
		Content: "You are an expert in various scientific disciplines, including physics, chemistry, and biology. Explain scientific concepts, theories, and phenomena in an engaging and accessible way. Use real-world examples and analogies to help users better understand and appreciate the wonders of science.",
		Sent:    time.Date(2025, 07, 03, 23, 50, 00, 00, time.UTC),
	},
	5: {
		ID:      5,
		Role:    types.RoleUser,
		Content: "Who are you?",
		Sent:    time.Date(2025, 07, 03, 23, 53, 00, 00, time.UTC),
	},
	6: {
		ID:      6,
		Role:    types.RoleAssistant,
		Content: "I'm an expert in various scientific disciplines, including physics, chemistry, and biology.",
		Sent:    time.Date(2025, 07, 03, 23, 55, 00, 00, time.UTC),
	},
}

func (s *MockStorage) ListChats() ([]types.Chat, error) {
	chats := make([]types.Chat, 2)
	for _, v := range allChats {
		chats = append(chats, v)
	}
	return chats, nil
}

func (s *MockStorage) GetChat(chatID int) (types.FullChat, error) {
	chat, ok := allChats[chatID]
	if !ok {
		return types.FullChat{}, &types.ErrChatNotFound{ChatID: chatID}
	}

	messages := make([]types.Message, 3)
	for _, messageID := range chat.MessageIDs {
		messages = append(messages, allMessages[messageID])
	}

	return types.FullChat{
		ID:        chat.ID,
		Title:     chat.Title,
		UpdatedAt: chat.UpdatedAt,
		Messages:  messages,
	}, nil
}

func (s *MockStorage) Init() error {
	return nil
}

func NewMockStorage() (*MockStorage, error) {
	return &MockStorage{}, nil
}
