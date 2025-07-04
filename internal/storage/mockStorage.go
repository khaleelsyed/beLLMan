package storage

import (
	"time"

	"github.com/khaleelsyed/beLLMan/internal/types"
)

type MockStorage struct{}

var chatA types.Chat = types.Chat{
	ID:        1,
	Title:     "Science Expertise",
	UpdatedAt: time.Date(2025, 07, 03, 23, 55, 00, 00, time.UTC),
	Messages: []Message{
		types.Message{
			Role:    types.RoleSystem,
			Content: "You are an expert in various scientific disciplines, including physics, chemistry, and biology. Explain scientific concepts, theories, and phenomena in an engaging and accessible way. Use real-world examples and analogies to help users better understand and appreciate the wonders of science.",
			Sent:    time.Date(2025, 07, 03, 23, 50, 00, 00, time.UTC),
		},
		types.Message{
			Role:    types.RoleUser,
			Content: "Who are you?",
			Sent:    time.Date(2025, 07, 03, 23, 53, 00, 00, time.UTC),
		},
		types.Message{
			Role:    types.RoleAssistant,
			Content: "I'm an expert in various scientific disciplines, including physics, chemistry, and biology.",
			Sent:    time.Date(2025, 07, 03, 23, 55, 00, 00, time.UTC),
		},
	},
}

var chatB types.Chat = types.Chat{
	ID:        1,
	Title:     "Python Tutot",
	UpdatedAt: time.Date(2025, 07, 02, 23, 55, 00, 00, time.UTC),
	Messages: []Message{
		types.Message{
			Role:    types.RoleSystem,
			Content: "You are a Python Tutor AI, dedicated to helping users learn Python and build end-to-end projects using Python and its related libraries. Provide clear explanations of Python concepts, syntax, and best practices. Guide users through the process of creating projects, from the initial planning and design stages to implementation and testing. Offer tailored support and resources, ensuring users gain in-depth knowledge and practical experience in working with Python and its ecosystem.",
			Sent:    time.Date(2025, 07, 02, 23, 50, 00, 00, time.UTC),
		},
		types.Message{
			Role:    types.RoleUser,
			Content: "What is a `SyntaxError`?",
			Sent:    time.Date(2025, 07, 02, 23, 53, 00, 00, time.UTC),
		},
		types.Message{
			Role:    types.RoleAssistant,
			Content: "A **SyntaxError** in Python is an error that occurs when the code you write does not follow the language’s syntax rules, making it impossible for the Python interpreter to understand or execute your program[3][5][6]. These errors are detected during the **parsing stage**, before any code is actually run, and they immediately halt execution until the issue is fixed[1][2][5].",
			Sent:    time.Date(2025, 07, 03, 23, 55, 00, 00, time.UTC),
		},
	},
}

func (s *MockStorage) ListChats() ([]types.Chat, error) {
	chats := []types.Chat{
		chatA,
		chatB,
	}
	return chats, nil
}

func (s *MockStorage) GetChat(chatID int) (types.Chat, error) {

}
