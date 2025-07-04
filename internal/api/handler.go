package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/khaleelsyed/beLLMan/internal/types"
)

var methodNotAllowed = func(w http.ResponseWriter) error {
	return WriteJSON(w, http.StatusMethodNotAllowed, APIError{"method not allowed"})
}

func (s *APIServer) ListChats(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return methodNotAllowed(w)
	}

	chats, err := s.storage.ListChats()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, chats)
}

func (s *APIServer) GetChat(w http.ResponseWriter, r *http.Request) error {
	var nf *types.ErrChatNotFound

	chatIDStr := mux.Vars(r)["id"]
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	chat, err := s.storage.GetChat(chatID)
	if err != nil {
		if errors.As(err, &nf) {
			return WriteJSON(w, http.StatusNotFound, nf.Error())
		}
		return WriteJSON(w, http.StatusInternalServerError, err)
	}

	return WriteJSON(w, http.StatusOK, chat)
}

func (s *APIServer) HandleChat(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.GetChat(w, r)
	default:
		return methodNotAllowed(w)
	}
}
