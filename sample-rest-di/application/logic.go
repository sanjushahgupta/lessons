package application

import (
	"github.com/sanjushahgupta/lessons/sample-rest-di/store/sqlite"
)

// Logic is struct that handles business logic in the application
type Logic struct {
	store *sqlite.Store
}

// NewLogic creates new instance of Logic
func NewLogic(store *sqlite.Store) *Logic {
	return &Logic{
		store: store,
	}
}
