package psql

import (
	provider "SQLbash/internal/db"
	"SQLbash/internal/domain"
	errorlib "SQLbash/internal/error"
	"context"
	"fmt"
)

type CommandRepository struct {
	db *provider.Provider
}

func NewCommandRepository(provider *provider.Provider) *CommandRepository {
	return &CommandRepository{
		db: provider,
	}
}

func (r *CommandRepository) ListCommands(ctx context.Context) ([]domain.Command, error) {
	var commands []domain.Command
	err := r.db.SelectContext(ctx, &commands, "SELECT * FROM commands")
	if err != nil {
		return nil, fmt.Errorf("failed to list commands: %w", errorlib.ErrHttpInternal)
	}
	return commands, nil
}

func (r *CommandRepository) GetCommand(ctx context.Context, id int) (*domain.Command, error) {
	var command domain.Command
	err := r.db.GetContext(ctx, &command, "SELECT * FROM commands WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get command: %w", errorlib.ErrHttpInvalidRequestData)
	}
	return &command, nil
}

func (r *CommandRepository) AddCommand(ctx context.Context, command string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO commands (command) VALUES ($1)", command)
	if err != nil {
		return fmt.Errorf("failed to add command: %w", errorlib.ErrHttpInternal)
	}
	return nil
}
