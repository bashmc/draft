package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Workspace represents a top-level organizational unit or collaboration space.
// Projects and Users belong to a Workspace.
type Workspace struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     uuid.UUID `json:"ownerId"`
	Owner       *User     `json:"owner,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type WorkspaceStore interface {
	Create(ctx context.Context, workspace *Workspace) error
	Update(ctx context.Context, workspace *Workspace) error
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (Workspace, error)
	GetAllForUser(ctx context.Context, userId uuid.UUID) ([]Workspace, error)
}
