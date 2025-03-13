package types

import (
	"github.com/google/uuid"
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	CreateUser(user *User) error
}

type CategoryStore interface {
	GetLatestAll() (*Category, error)
	GetById(id uuid.UUID) (*Category, error)
	GetBySlug(slug string) (*Category, error)
	Create(category *Category) error
	Update(id uuid.UUID, category *Category) (*Category, error)
	Delete(id uuid.UUID) (error, bool)
}

type Category struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateCategoryPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required,min=4"`
	LastName  string `json:"last_name"  validate:"required,min=4"`
	Email     string `json:"email"     validate:"required,email"`
	Password  string `json:"password"  validate:"required,min=6"`
}

type LoginUserPayload struct {
	Email    string `json:"email"     validate:"required,email"`
	Password string `json:"password"  validate:"required,min=6"`
}
