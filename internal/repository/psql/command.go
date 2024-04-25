package psql

import (
	provider "bashExecAPI/internal/db"
	"bashExecAPI/internal/domain"
	errorlib "bashExecAPI/internal/error"
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

func (r *CommandRepository) GetCommands(ctx context.Context) ([]domain.Command, error) {
	var commands []domain.Command
	err := r.db.SelectContext(ctx, &commands, "SELECT * FROM commands")
	if err != nil {
		return nil, fmt.Errorf("failed to list commands: %w", errorlib.ErrHttpInternal)
	}
	return commands, nil
}

func (r *CommandRepository) GetCommand(ctx context.Context, id string) (*domain.Command, error) {
	var command domain.Command
	err := r.db.GetContext(ctx, &command, "SELECT * FROM commands WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get command: %w", errorlib.ErrHttpInvalidRequestData)
	}
	return &command, nil
}

func (r *CommandRepository) CreateCommand(ctx context.Context, command string) (string, error) {
	// Implement logic to insert a new command into the database and return its ID
	return "", fmt.Errorf("not implemented")
}

func (r *CommandRepository) StopCommand(ctx context.Context, id string) error {
	// Implement logic to stop a running command by ID
	return fmt.Errorf("not implemented")
}
