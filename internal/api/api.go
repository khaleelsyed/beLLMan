package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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
	router.HandleFunc("/chats", makeHTTPHandlerFunc(s.ListChats))
	router.HandleFunc("/chat/{id}", makeHTTPHandlerFunc(s.HandleChat))

	allowedOrigins := handlers.AllowedOrigins([]string{os.Getenv("CORS_ALLOWED_SINGLE_ORIGIN")})
	allowedMethods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	router.Use(handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders))

	log.Println("Listening for connections on ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
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

	err, errFound := v.(error)
	if errFound && status < 400 {
		status = http.StatusInternalServerError
	}

	w.WriteHeader(status)

	// TODO: See if this is necessary, and if the below return will suffice?
	// Is there actually any difference in the output between this and the latter?
	if errFound {
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
