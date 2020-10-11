package domain

// User is a struct for user related operations
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
}
