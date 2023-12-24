package repository

import (
	"database/sql"
	"echobot/types"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

func NewSqlite() (*Sqlite, error) {
	db, err := sql.Open("sqlite3", "data/sqlite/bot.db")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Sqlite{db}, nil
}

func (s *Sqlite) CreateUser(user types.User) error {
	q := "INSERT INTO users (chat_id, region_id, username, lang) VALUES(?, ?, ?, ?)"
	_, err := s.db.Exec(q, user.ChatID, user.RegionID, user.Username, user.Language)
	return err
}

func (s *Sqlite) GetRegionID(chatID int64) (int, error) {
	q := "SELECT region_id FROM users WHERE chat_id = ?"
	r := s.db.QueryRow(q, chatID)
	var regionID int
	err := r.Scan(&regionID)
	return regionID, err
}

func (s *Sqlite) GetAllUsersByRegionID(regionID int) ([]int64, error) {
	q := "SELECT chat_id FROM users WHERE region_id = ?"
	r, err := s.db.Query(q, regionID)
	if err != nil {
		return nil, err
	}
	var chatIDs []int64
	for r.Next() {
		var chatID int64
		if err := r.Scan(&chatID); err != nil {
			return nil, err
		}
		chatIDs = append(chatIDs, chatID)
	}
	return chatIDs, r.Err()
}

func (s *Sqlite) Init() error {
	q := "CREATE TABLE IF NOT EXISTS users(chat_id INTEGER UNIQUE, region_id INTEGER, lang TEXT DEFAULT 'tj', last_message_id INTEGER DEFAULT 0, username text)"
	_, err := s.db.Exec(q)
	return err
}

func (s *Sqlite) UpdateRegionID(chatID int64, regionID int) error {
	q := "UPDATE users SET region_id = ? WHERE chat_id = ?"
	_, err := s.db.Exec(q, regionID, chatID)
	return err
}

func (s *Sqlite) UpdateLanguage(chatID int64, language string) error {
	q := "UPDATE users SET lang = ? WHERE chat_id = ?"
	_, err := s.db.Exec(q, language, chatID)
	return err
}

func (s *Sqlite) GetLang(chatID int64) (string, error) {
	q := "SELECT lang FROM users WHERE chat_id = ?"
	r := s.db.QueryRow(q, chatID)
	var lang string
	err := r.Scan(&lang)
	return lang, err
}

func (s *Sqlite) UpdateLastMessageID(chatID int64, lastMessageID int) error {
	q := "UPDATE users SET last_message_id = ? WHERE chat_id = ?"
	_, err := s.db.Exec(q, lastMessageID, chatID)
	return err
}

func (s *Sqlite) GetLastMessageID(chatID int64) (int, error) {
	q := "SELECT last_message_id FROM users WHERE chat_id = ?"
	r := s.db.QueryRow(q, chatID)
	var last_msg_id int
	err := r.Scan(&last_msg_id)
	return last_msg_id, err
}

func (s *Sqlite) DeleteUser(chatID int64) error {
	q := "DELETE FROM users WHERE chat_id = ?"
	_, err := s.db.Exec(q, chatID)
	return err
}
