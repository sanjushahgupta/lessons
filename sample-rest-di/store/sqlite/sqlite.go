package sqlite

import (
	"database/sql"
)

// Store is struct for handling SQLite database operations
type Store struct {
	db *sql.DB
}

// NewStore creates new instance of Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
