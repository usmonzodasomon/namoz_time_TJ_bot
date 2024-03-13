package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	User     string
	DBName   string
	SSLMode  string
	Password string
	Host     string
}

func GetDBConnection(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s",
		cfg.User, cfg.DBName, cfg.SSLMode, cfg.Password, cfg.Host)

	return sqlx.Connect("postgres", dsn)
}

func CloseDBConnection(dbConn *sqlx.DB) error {
	return dbConn.Close()
}
