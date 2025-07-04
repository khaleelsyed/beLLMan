package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khaleelsyed/beLLMan/internal/storage"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error
type APIError struct{ Error string }

type APIServer struct {
	listenAddr string
	storage    storage.Storage
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/chats", makeHTTPHandlerFunc())
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err, errFound := v.(error); errFound {
		return json.NewEncoder(w).Encode(err)
	}

	return json.NewEncoder(w).Encode(v)
}

func NewAPIServer(listenAddr string, storage storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    storage,
	}
}
