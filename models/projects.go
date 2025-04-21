package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID  `json:"id"`
	WorkspaceID uuid.UUID  `json:"workspaceId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	OwnerID     uuid.UUID  `json:"ownerId"`
	Owner       *User      `json:"owner,omitempty"`
	Workspace   *Workspace `json:"workspace,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
