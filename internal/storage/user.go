package storage

import (
	"context"
	"simpleserver/internal/models"
)

func (s *Storage) CreateUser(ctx context.Context, email string, passwordHash string) (models.User, error) {
	var user models.User

	err := s.DB.QueryRow(
		ctx,
		`
		INSERT INTO users (email, password_hash, role)
		VALUES ($1, $2, $3)
		RETURNING id, email, password_hash, role
		`,
		email,
		passwordHash,
		models.RoleUser,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role)

	return user, err
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	err := s.DB.QueryRow(
		ctx,
		`
		SELECT id, email, password_hash, role
		FROM users
		WHERE email = $1
		`,
		email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role)

	return user, err
}
