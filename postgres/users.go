package postgres

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/bashmc/draft/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore struct {
	conn *pgxpool.Pool
}

func NewUserStore(conn *pgxpool.Pool) models.UserStore {
	return &UserStore{
		conn: conn,
	}
}

// InsertUser implements models.UserStore.
func (u *UserStore) InsertUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (id, name, email, password, profile_photo, created_at, updated_at, verified)
		VALUES ($1, NULLIF($2,''), $3, $4, $5, $6, $7, $8);`

	_, err := u.conn.Exec(ctx, query,
		user.Id,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.ProfilePhoto,
		user.CreatedAt,
		user.UpdatedAt,
		user.Verified,
	)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return models.ErrDuplicateUser
		}
		slog.Error("failed to insert user", "error", err)
		return err
	}
	return nil
}

// DeleteUser implements models.UserStore.
func (u *UserStore) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1;`

	result, err := u.conn.Exec(ctx, query, id)
	if err != nil {
		slog.Error("failed delete user", "error", err)
		return err
	}

	if result.RowsAffected() == 0 {
		return models.ErrUserNotFound
	}
	return nil
}

// GetUser implements models.UserStore.
func (u *UserStore) GetUser(ctx context.Context, id uuid.UUID) (models.User, error) {
	query := `
		SELECT id, name, email, password, profile_photo, created_at, updated_at, verified 
		FROM users 
		WHERE id = $1;`

	var user models.User
	err := u.conn.QueryRow(ctx, query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.ProfilePhoto,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Verified,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return models.User{}, models.ErrUserNotFound
	}

	return user, nil
}

// UpdateUser implements models.UserStore.
func (u *UserStore) UpdateUser(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET name = $1, email = $2, password = $3, profile_photo = $4, updated_at = $5, verified = $6
		WHERE id = $7;`

	result, err := u.conn.Exec(ctx, query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.ProfilePhoto,
		user.UpdatedAt,
		user.Verified,
		user.Id,
	)
	if err != nil {
		slog.Error("failed update user", "error", err)
		return err
	}

	if result.RowsAffected() == 0 {
		return models.ErrUserNotFound
	}

	return nil
}
