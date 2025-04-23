-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workspaces(
    id uuid NOT NULL,
    name VARCHAR(120) NOT NULL,
    description TEXT,
    ownerId uuid NOT NULL,
    createdAt TIMESTAMP DEFAULT now() NOT NULL,
    updatedAt TIMESTAMP DEFAULT now() NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(ownerId) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workspaces;
-- +goose StatementEnd
