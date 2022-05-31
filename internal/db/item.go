package db

import (
	"database/sql"

	"github.com/karolsudol/fiesta/internal/models"
)

func (db Database) GetAllItems() (*models.ItemList, error) {
	list := &models.ItemList{}

	rows, err := db.Conn.Query("SELECT * FROM items ORDER BY ID DESC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.City)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

func (db Database) AddItem(item *models.Item) error {
	var id int
	var createdAt string
	query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, item.Name).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	item.ID = id
	// item.CreatedAt = createdAt
	return nil
}

func (db Database) GetItemById(itemId int) (models.Item, error) {
	item := models.Item{}

	query := `SELECT * FROM items WHERE id = $1;`
	row := db.Conn.QueryRow(query, itemId)
	switch err := row.Scan(&item.ID, &item.Name); err {
	case sql.ErrNoRows:
		return item, ErrNoMatch
	default:
		return item, err
	}
}

func (db Database) DeleteItem(itemId int) error {
	query := `DELETE FROM items WHERE id = $1;`
	_, err := db.Conn.Exec(query, itemId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
	item := models.Item{}
	query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
	err := db.Conn.QueryRow(query, itemData.Name, itemId).Scan(&item.ID, &item.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, ErrNoMatch
		}
		return item, err
	}

	return item, nil
}
