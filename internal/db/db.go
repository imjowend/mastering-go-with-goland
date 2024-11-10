package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Item struct {
	Task   string
	Status string
}

type DB struct {
	pool *pgxpool.Pool
}

func New(user, password, host, dbname string, port int) (*DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the db: %w", err)
	}
	return &DB{pool: pool}, nil
}

func (db *DB) InsertItem(ctx context.Context, item Item) error {
	// for query use: `.
	query := `INSERT INTO todo_items (task, status) VALUES ($1, $2)`
	_, err := db.pool.Exec(ctx, query, item.Task, item.Status)
	return err
}

func (db *DB) GetAllItems(ctx context.Context) ([]Item, error) {
	query := `SELECT task, status FROM todo_items`
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Task, &item.Status); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (db *DB) Close() {
	db.pool.Close()
}
