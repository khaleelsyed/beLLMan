package api

func (s *APIServer) ListChats() {
	chats := s.storage.ListChats()
}
