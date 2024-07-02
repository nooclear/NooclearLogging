package NooclearLogging

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDB(dbPath string) error {
	var err error
	if db, err = sql.Open("sqlite", dbPath); err != nil {
		return err
	}
	if _, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS logs (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	dateTime TEXT NOT NULL,
    	category TEXT NOT NULL,
    	message TEXT NOT NULL
		)`,
	); err != nil {
		return err
	}

	return nil
}

func addLog(l *Log) (int64, error) {
	result, err := db.ExecContext(
		context.Background(),
		`INSERT INTO logs (dateTime, category, message) VALUES (?, ?, ?)`, l.Datetime, l.Category, l.Message,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
