package database

import (
	"database/sql"
	"todolist/internal/entity"
	uuid "todolist/pkg/entity"
)

type ListDB struct {
	DB *sql.DB
}

func NewListDB(db *sql.DB) *ListDB {
	return &ListDB{DB: db}
}

func (u *ListDB) Create(list *entity.List) error {
	err := u.DB.QueryRow(`INSERT INTO lists (id, title, description, status) VALUES ($1, $2, $3, $4);`, list.ID, list.Title, list.Description, list.Status).Err()

	if err != nil {
		return err
	}

	return nil
}

func (u *ListDB) Delete(id uuid.ID) error {
	err := u.DB.QueryRow(`DELETE FROM lists WHERE id = $1;`, id).Err()

	if err != nil {
		return err
	}

	return nil
}

func (u *ListDB) FindById(id uuid.ID) (*entity.List, error) {
	list := entity.List{}

	row := u.DB.QueryRow(`SELECT id, title, description, status, created_at, updated_at FROM lists WHERE id = $1;`, id)

	if err := row.Scan(&list.ID, &list.Title, &list.Description, &list.Status, &list.CreatedAt, &list.UpdatedAt); err != nil {
		return nil, err
	}

	return &list, nil
}
