package models

import (
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
