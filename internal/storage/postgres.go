package storage

import (
	"context"
	"simpleserver/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	DB *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (*Storage, error) {
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) Close() {
	s.DB.Close()
}

func (s *Storage) CreatePanel(ctx context.Context, title string, description string) error {
	_, err := s.DB.Exec(
		ctx,
		"INSERT INTO panels (title, description) VALUES ($1, $2)",
		title,
		description,
	)
	return err
}

func (s *Storage) GetPanels(ctx context.Context) ([]models.Panel, error) {
	rows, err := s.DB.Query(
		ctx,
		"SELECT id, title, description FROM panels ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var panels []models.Panel
	for rows.Next() {
		var panel models.Panel
		err := rows.Scan(&panel.ID, &panel.Title, &panel.Description)
		if err != nil {
			return nil, err
		}
		panels = append(panels, panel)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return panels, nil
}
