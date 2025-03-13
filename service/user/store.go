package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
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
	row := s.db.QueryRow(`SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = ?`, email)

	var user types.User
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func scanRowIntroUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) CreateUser(user *types.User) error {
	query := `INSERT INTO users (email, password, first_name, last_name, id) VALUES (?, ?, ?, ?, ?)`
	if _, err := s.db.Exec(query, user.Email, user.Password, user.FirstName, user.LastName, user.ID); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByID(id uuid.UUID) (*types.User, error) {
	row := s.db.QueryRow(`SELECT id, first_name, last_name, email, created_at FROM users WHERE id = ?`, id)

	var user types.User
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
