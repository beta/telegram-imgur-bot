package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq" // For connecting to PostgreSQL DB.
)

var (
	db       *sql.DB
	initOnce = new(sync.Once)
)

// Init initializes database connection.
func Init(dsn string) error {
	var err error
	initOnce.Do(func() {
		conn, connErr := sql.Open("postgres", dsn)
		if connErr != nil {
			err = fmt.Errorf("error while connecting to DB: %v", connErr)
			return
		}

		if pingErr := conn.Ping(); pingErr != nil {
			err = fmt.Errorf("error while checking DB connection: %v", pingErr)
			return
		}
		db = conn
	})

	return err
}
