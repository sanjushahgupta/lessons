package sqlite

import (
	"github.com/sanjushahgupta/lessons/sample-rest-di/domain"
)

// CreateUser creates user
func (s *Store) CreateUser(user domain.User) (int, error) {
	var id int

	tx, err := s.db.Begin()
	if err != nil {
		return id, err
	}

	stmt, err := tx.Prepare("insert into user(first_name, last_name, email) values(?, ?, ?)")
	if err != nil {
		return id, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email)
	if err != nil {
		return id, err
	}

	err = tx.Commit()
	if err != nil {
		return id, err
	}

	id64, _ := res.LastInsertId()
	return int(id64), nil
}

// GetAllUser gets all user
func (s *Store) GetAllUser() ([]domain.User, error) {
	var users []domain.User

	rows, err := s.db.Query("select id, first_name, last_name, email from user")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName, lastName, email string
		err = rows.Scan(&id, &firstName, &lastName, &email)
		if err != nil {
			return users, err
		}

		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}
	return users, nil
}

// GetUser gets specific user by ID
// TODO allow generic query instead of by ID
// DETAILS generic query (HARD)
func (s *Store) GetUser(id int) (domain.User, error) {
	var user domain.User
	rows, err := s.db.Prepare("select first_name, last_name, email from user where id = ?")
	if err != nil {
		return user, err
	}
	var firstName, lastName, email string
	err = rows.QueryRow(id).Scan(&firstName, &lastName, &email)
	if err != nil {
		return user, err
	}

	user = domain.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	return user, nil
}
