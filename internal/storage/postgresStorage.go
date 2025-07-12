package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/khaleelsyed/beLLMan/internal/types"
)

type PostresStorage struct {
	db *sql.DB
}

func (s *PostresStorage) GetChat(chatID int) (types.Chat, error) {
	result, err := s.db.Query(`SELECT id, title, updated_at FROM chat WHERE id = $1`, chatID)
	if err != nil {
		return types.Chat{}, err
	}

	defer result.Close()

	for result.Next() {
		var chat types.Chat

		err := result.Scan(&chat.ID, &chat.Title, &chat.UpdatedAt)
		if err != nil {
			return types.Chat{}, nil
		} else {
			return chat, nil
		}
	}

	return types.Chat{}, &types.ErrChatNotFound{ChatID: chatID}
}

func (s *PostresStorage) ListChats() ([]types.Chat, error) {
	rows, err := s.db.Query(`SELECT id, title, updated_at FROM chat`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	chatsOutput := make([]types.Chat, 0)
	for rows.Next() {
		var chat types.Chat
		err := rows.Scan(&chat.ID, &chat.Title, &chat.UpdatedAt)
		if err != nil {
			return nil, err
		}

		chatsOutput = append(chatsOutput, chat)
	}

	if len(chatsOutput) == 0 {
		return nil, errors.New("no chats found")
	}

	return chatsOutput, nil
}

func NewPostgresStorage() (*PostresStorage, error) {

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s port=%s sslmode=%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSL_MODE"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostresStorage{db: db}, nil
}

func (s *PostresStorage) Init() error {
	var err error
	if err = s.createRoleTable(); err != nil {
		return err
	}

	if err = s.createChatTable(); err != nil {
		return err
	}

	if err = s.createMessageTable(); err != nil {
		return err
	}

	if err = s.createChatUpdatedFunction(); err != nil {
		return err
	}

	return nil
}

func (s *PostresStorage) createRoleTable() error {
	query := `CREATE TABLE IF NOT EXISTS role(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL UNIQUE
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostresStorage) createChatTable() error {
	query := `CREATE TABLE IF NOT EXISTS chat(
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostresStorage) createMessageTable() error {
	query := `CREATE TABLE IF NOT EXISTS message(
	id SERIAL PRIMARY KEY,
	role SERIAL,
	CONSTRAINT fk_role FOREIGN KEY(role) REFERENCES role(id) ON DELETE RESTRICT,
	content TEXT NOT NULL,
	chat_id SERIAL,
	CONSTRAINT fk_chat FOREIGN KEY(chat_id) REFERENCES chat(id) ON DELETE CASCADE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostresStorage) createChatUpdatedFunction() error {
	query := `CREATE OR REPLACE FUNCTION update_chat_updated_at()
	RETURNS TRIGGER AS $$
	BEGIN
	  IF (TG_OP = 'INSERT') THEN
	    UPDATE public.chat
	    SET updated_at = now()
	    WHERE id = NEW.chat_id;

	  ELSIF (TG_OP = 'UPDATE') THEN
	    IF (OLD.chat_id <> NEW.chat_id) THEN
	      UPDATE public.chat
	      SET updated_at = now()
	      WHERE id = OLD.chat_id;
	    END IF;

	    UPDATE public.chat
	    SET updated_at = now()
	    WHERE id = NEW.chat_id;

	  ELSIF (TG_OP = 'DELETE') THEN
	    UPDATE public.chat
	    SET updated_at = now()
	    WHERE id = OLD.chat_id;
	  END IF;

	  RETURN NULL;
	END;
	$$ LANGUAGE plpgsql;
	
	DROP TRIGGER IF EXISTS trigger_update_chat_updated_at ON public.message;

	CREATE TRIGGER trigger_update_chat_updated_at
	AFTER INSERT OR UPDATE OR DELETE ON public.message
	FOR EACH ROW
	EXECUTE FUNCTION update_chat_updated_at();
	`

	_, err := s.db.Exec(query)
	return err
}
