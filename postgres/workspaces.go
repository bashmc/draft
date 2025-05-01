package postgres

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/topbash/draft/models"
)

type WorkspaceStore struct {
	conn *pgxpool.Pool
}

func NewWorkspaceStore(conn *pgxpool.Pool) models.WorkspaceStore {
	return &WorkspaceStore{conn: conn}
}

// Create implements models.WorkspaceStore.
func (w *WorkspaceStore) Create(ctx context.Context, workspace *models.Workspace) error {
	query := `INSERT INTO workspaces(id, name, description, user_id, created_at, updated_at)
	VALUE($1, $2, $3, $4, now(), now());`

	_, err := w.conn.Exec(ctx, query, &workspace.CreatedAt, &workspace.Name, &workspace.Description, &workspace.OwnerID)
	if err != nil {
		slog.Error("failed to create workspace", "error", err)
		return err
	}

	return nil
}

// Delete implements models.WorkspaceStore.
func (w *WorkspaceStore) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM workspaces WHERE id = $1;`

	_, err := w.conn.Exec(ctx, query, id)

	if err != nil {
		slog.Error("failed to create workspace", "error", err)
		return err
	}

	return nil
}

// Get implements models.WorkspaceStore.
func (w *WorkspaceStore) Get(ctx context.Context, id uuid.UUID) (models.Workspace, error) {
	query := `SELECT id, name, descrption, user_id, created_at, updated_at
	FROM workspaces WHERE id = $1;`

	row := w.conn.QueryRow(ctx, query, id)

	ws := models.Workspace{}
	err := row.Scan(&ws.Id, &ws.Name, &ws.Description, &ws.OwnerID, &ws.CreatedAt, &ws.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Workspace{}, models.ErrNotFound
		}
		slog.Error("failed to read workspace", "error", err)
		return models.Workspace{}, err
	}

	return ws, nil
}

// GetAllForUser implements models.WorkspaceStore.
func (w *WorkspaceStore) GetAllForUser(ctx context.Context, userId uuid.UUID) ([]models.Workspace, error) {
	panic("unimplemented")
}

// Update implements models.WorkspaceStore.
func (w *WorkspaceStore) Update(ctx context.Context, workspace *models.Workspace) error {
	panic("unimplemented")
}
