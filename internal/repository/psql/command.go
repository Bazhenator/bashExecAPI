package psql

import (
	provider "bashExecAPI/internal/db"
	"bashExecAPI/internal/domain"
	errorlib "bashExecAPI/internal/error"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
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
	query := `
		INSERT INTO commands (command) VALUES ($1) RETURNING id;
	`

	var id int64
	err := r.db.QueryRowContext(ctx, query, command).Scan(&id)
	if err != nil {
		log.Error(fmt.Errorf("failed to insert command into database: %v", err))
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}

func (r *CommandRepository) RunCommand(ctx context.Context, id string) (string, error) {
	command, err := r.GetCommand(ctx, id)
	if err != nil {
		log.Error(fmt.Errorf("failed to get command: %v", err))
		return "", err
	}
	file, err := os.Create("script.sh")
	if err != nil {
		log.Error(fmt.Errorf("failed to open script file: %v", err))
		return "", err
	}
	defer os.Remove(file.Name())

	if _, err := file.WriteString("#!/bin/bash\n" + command.Command); err != nil {
		log.Error(fmt.Errorf("failed to write command to script file: %v", err))
		return "", err
	}
	if err := file.Close(); err != nil {
		log.Error(fmt.Errorf("failed to close script file: %v", err))
		return "", err
	}

	if err := os.Chmod(file.Name(), 0755); err != nil {
		log.Error(fmt.Errorf("failed to set execute permissions on script file: %v", err))
		return "", err
	}

	cmd := exec.Command("bash", file.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Errorf("failed to execute command: %v", err))
		return "", err
	}

	return string(output), nil
}

func (r *CommandRepository) StopCommand(ctx context.Context, id string) error {

	return fmt.Errorf("not implemented")
}
