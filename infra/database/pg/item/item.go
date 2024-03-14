package item

import (
	"database/sql"
	"todolist/internal/entity/item"
	uuid "todolist/pkg/entity"
)

type ItemDB struct {
	DB *sql.DB
}

func NewItemDB(db *sql.DB) *ItemDB {
	return &ItemDB{DB: db}
}

func (u *ItemDB) Create(item *item.Item) error {
	err := u.DB.QueryRow(
		`INSERT INTO items (id, title, description, status, list_id) VALUES ($1, $2, $3, $4, $5);`,
		item.ID, item.Title, item.Description, item.Status, item.ListId).Err()

	if err != nil {
		return err
	}

	return nil
}

func (u *ItemDB) Delete(id uuid.ID) error {
	err := u.DB.QueryRow(`DELETE FROM items WHERE id = $1;`, id).Err()

	if err != nil {
		return err
	}

	return nil
}

func (u *ItemDB) FindById(id uuid.ID) (*item.Item, error) {
	item := item.Item{}

	row := u.DB.QueryRow(`SELECT id, title, description, status FROM items WHERE id = $1;`, id)

	if err := row.Scan(&item.ID, &item.Title, &item.Description, &item.Status); err != nil {
		return nil, err
	}

	return &item, nil
}
