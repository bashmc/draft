package models

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrDuplicateUser = errors.New("user with email already exists")
)

type User struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Active       bool      `json:"active"`
}

type UserStore interface {
	InsertUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, Id uuid.UUID) (User, error)
	DeleteUser(ctx context.Context, Id string) error
}
