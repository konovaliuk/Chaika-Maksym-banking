package db

import (
	"database/sql"
)

type Cleaner struct {
	db *sql.DB
}

func NewCleaner(db *sql.DB) *Cleaner {
	return &Cleaner{db: db}
}

func (c *Cleaner) CleanDatabase() error {
	rows, err := c.db.Query(
		"SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_name NOT LIKE '%migration%'",
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return err
		}

		_, err := c.db.Exec("DELETE FROM " + tableName)
		if err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
