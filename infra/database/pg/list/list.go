package list

import (
	"database/sql"
	"todolist/internal/entity/list"
	uuid "todolist/pkg/entity"
)

type ListDB struct {
	DB *sql.DB
}

func NewListDB(db *sql.DB) *ListDB {
	return &ListDB{DB: db}
}

func (u *ListDB) Create(list *list.List) error {
	err := u.DB.QueryRow(
		`INSERT INTO lists (id, title, description, status, owner_id) VALUES ($1, $2, $3, $4, $5);`,
		list.ID, list.Title, list.Description, list.Status, list.Owner.ID.String()).Err()

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

func (u *ListDB) FindById(id uuid.ID) (*list.List, error) {
	list := list.List{}

	row := u.DB.QueryRow(`SELECT id, title, description, status, created_at, updated_at FROM lists WHERE id = $1;`, id)

	if err := row.Scan(&list.ID, &list.Title, &list.Description, &list.Status, &list.CreatedAt, &list.UpdatedAt); err != nil {
		return nil, err
	}

	return &list, nil
}
