package user

import (
	"database/sql"
	"fmt"
	"github.com/pimp13/go-react-project/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

var _ types.UserStore = (*Store)(nil)

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT `id`, `email` FROM `users` WHERE `email` = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user *types.User
	if rows.Next() {
		user, err = scanRowIntroUser(rows)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func scanRowIntroUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(&user.ID, &user.Email)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) CreateUser(user *types.User) error {
	query := `INSERT INTO users (email, password, firstName, lastName) VALUES (?, ?, ?, ?)`
	_, err := s.db.Exec(query, user.Email, user.Password, user.FirstName, user.LastName)
	return err
}

func (s *Store) GetUserByID(id uint) (*types.User, error) {
	return nil, nil
}
