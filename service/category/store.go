package category

import (
	"database/sql"
	"github.com/google/uuid"

	"github.com/pimp13/go-react-project/types"
)

type Store struct {
	db *sql.DB
}

var _ types.CategoryStore = (*Store)(nil)

func (s *Store) GetLatestAll() (*types.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GetById(id uuid.UUID) (*types.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GetBySlug(slug string) (*types.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) Create(category *types.Category) error {
	query := `
		INSERT INTO categories (id, name, description, slug, image, created_at) 
		VALUES (?, ?, ?, ?, ?, ?)
		`
	_, err := s.db.Exec(query, category.ID, category.Name, category.Description, category.Slug, category.Image, category.CreatedAt)
	return err
}

func (s *Store) Update(id uuid.UUID, category *types.Category) (*types.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) Delete(id uuid.UUID) (error, bool) {
	//TODO implement me
	panic("implement me")
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
