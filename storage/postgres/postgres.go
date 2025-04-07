package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
	"strings"
)

type Storage struct {
	db *sqlx.DB
}

func NewPostgresStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) AddUserIfNotExist(user types.User) error {
	q := `SELECT is_deleted from users where chat_id = $1`
	var isDeleted bool
	err := s.db.Get(&isDeleted, q, user.ChatID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			q = `INSERT INTO users(chat_id, region_id, username, lang) VALUES($1, $2, $3, $4)`
			_, err = s.db.Exec(q, user.ChatID, user.RegionID, user.Username, user.Language)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	q = `UPDATE users SET is_deleted = false WHERE chat_id = $1`
	_, err = s.db.Exec(q, user.ChatID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetUser(chatID int64) (types.User, error) {
	q := "SELECT * FROM users WHERE chat_id = $1"
	var user types.User
	user.Language = "tj"
	err := s.db.Get(&user, q, chatID)
	return user, err
}

func (s *Storage) UpdateUser(user types.User) error {
	q := `UPDATE users SET`
	queries := make([]string, 0)
	params := make([]interface{}, 0)

	placeHolderCount := 1
	if user.RegionID != 0 {
		queries = append(queries, fmt.Sprintf(" region_id = $%d", placeHolderCount))
		params = append(params, user.RegionID)
		placeHolderCount++
	}

	if user.Language != "" {
		queries = append(queries, fmt.Sprintf(" lang = $%d", placeHolderCount))
		params = append(params, user.Language)
		placeHolderCount++
	}

	if user.LastMessageID != 0 {
		queries = append(queries, fmt.Sprintf(" last_message_id = $%d", placeHolderCount))
		params = append(params, user.LastMessageID)
		placeHolderCount++
	}

	q += strings.Join(queries, ",")

	q += fmt.Sprintf(" WHERE chat_id = $%d", placeHolderCount)
	params = append(params, user.ChatID)
	_, err := s.db.Exec(q, params...)
	return err
}

func (s *Storage) DeleteUser(chatID int64) error {
	q := "UPDATE users SET is_deleted = true WHERE chat_id = $1"
	_, err := s.db.Exec(q, chatID)
	return err
}

func (s *Storage) GetAllUsersByRegionID(regionID int) ([]types.User, error) {
	q := "SELECT * FROM users WHERE is_deleted = false AND region_id = $1"
	var users []types.User
	err := s.db.Select(&users, q, regionID)
	return users, err
}

func (s *Storage) GetNamazTime(date string) (types.NamazTime, error) {
	q := "SELECT * FROM namaz_time WHERE date = $1"
	var namazTime types.NamazTime
	err := s.db.Get(&namazTime, q, date)
	return namazTime, err
}

func (s *Storage) GetRegionByID(lang string, id int) (types.Region, error) {
	var region types.Region
	switch lang {
	case "ru":
		q := `SELECT * FROM ru_regions WHERE id = $1`
		if err := s.db.Get(&region, q, id); err != nil {
			return types.Region{}, err
		}
		return region, nil
	case "tj":
		q := `SELECT * FROM tj_regions WHERE id = $1`
		if err := s.db.Get(&region, q, id); err != nil {
			return types.Region{}, err
		}
		return region, nil
	}
	return types.Region{}, nil
}

func (s *Storage) UpdateNamazTime(namazTime []types.NamazTime) error {
	q := `DELETE FROM namaz_time WHERE 1 = 1`
	_, err := s.db.Exec(q)
	if err != nil {
		return err
	}

	q = `INSERT INTO namaz_time(date, fajr_from, fajr_to, zuhr_from, zuhr_to, asr_from, asr_to, maghrib_from, maghrib_to, isha_from, isha_to) 
		VALUES(:date, :fajr_from, :fajr_to, :zuhr_from, :zuhr_to, :asr_from,
				:asr_to, :maghrib_from, :maghrib_to, :isha_from, :isha_to)`

	_, err = s.db.NamedExec(q, namazTime)
	if err != nil {
		return err
	}
	return nil
}

//
//func (s *Storage) GetRegionID(chatID int64) (int, error) {
//	q := "SELECT region_id FROM users WHERE chat_id = ?"
//	r := s.db.QueryRow(q, chatID)
//	var regionID int
//	err := r.Scan(&regionID)
//	return regionID, err
//}

//
//func (s *Sqlite) Init() error {
//	q := `
//CREATE TABLE IF NOT EXISTS users(
//    chat_id INTEGER UNIQUE,
//    region_id INTEGER,
//    lang TEXT DEFAULT 'tj',
//    last_message_id INTEGER DEFAULT 0,
//     username text,
//     is_deleted BOOLEAN DEFAULT false)`
//	_, err := s.db.Exec(q)
//	return err
//}
//
//func (s *Sqlite) UpdateRegionID(chatID int64, regionID int) error {
//	q := "UPDATE users SET region_id = ? WHERE chat_id = ?"
//	_, err := s.db.Exec(q, regionID, chatID)
//	return err
//}
//
//func (s *Sqlite) UpdateLanguage(chatID int64, language string) error {
//	q := "UPDATE users SET lang = ? WHERE chat_id = ?"
//	_, err := s.db.Exec(q, language, chatID)
//	return err
//}
//
//func (s *Sqlite) GetLang(chatID int64) (string, error) {
//	q := "SELECT lang FROM users WHERE chat_id = ?"
//	r := s.db.QueryRow(q, chatID)
//	var lang string
//	err := r.Scan(&lang)
//	if lang == "" {
//		lang = "tj"
//	}
//	return lang, err
//}
//
//func (s *Sqlite) UpdateLastMessageID(chatID int64, lastMessageID int) error {
//	q := "UPDATE users SET last_message_id = ? WHERE chat_id = ?"
//	_, err := s.db.Exec(q, lastMessageID, chatID)
//	return err
//}
//
//func (s *Sqlite) GetLastMessageID(chatID int64) (int, error) {
//	q := "SELECT last_message_id FROM users WHERE chat_id = ?"
//	r := s.db.QueryRow(q, chatID)
//	var LastMsgID int
//	err := r.Scan(&LastMsgID)
//	return LastMsgID, err
//}
