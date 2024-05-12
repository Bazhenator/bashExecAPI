package psql

import (
	"context"
	"database/sql"
	"fmt"
	provider "github.com/Bazhenator/bashExecAPI/internal/db"
	"github.com/Bazhenator/bashExecAPI/internal/domain"
	errorlib "github.com/Bazhenator/bashExecAPI/internal/error"
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

func (r *CommandRepository) CreateCommand(ctx context.Context, command string) (string, string, error) {
	query := `
		INSERT INTO commands (command) VALUES ($1) RETURNING id;
	`

	var id int
	err := r.db.QueryRowContext(ctx, query, command).Scan(&id)
	if err != nil {
		log.Error(fmt.Errorf("failed to insert command into database: %v", err))
		return "", "", err
	}

	var result string
	result, err = r.RunCommand(ctx, id)

	query = `
        UPDATE commands SET result = $1 WHERE id = $2;
    `
	_, err = r.db.ExecContext(ctx, query, result, id)
	if err != nil {
		log.Error(fmt.Errorf("failed to update command result in database: %v", err))
		return "", "", err
	}
	return result, fmt.Sprintf("%d", id), nil
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
	err := r.db.GetContext(ctx, &command, "SELECT id, command, result FROM commands WHERE id = $1", id)
	if err != nil {
		log.Error(fmt.Errorf("failed to get command from database: %v", err))
		if err == sql.ErrNoRows {
			log.Error(fmt.Errorf("command with id %d not found", id))
			return nil, fmt.Errorf("command with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to get command: %w", err)
	}

	return &command, nil
}

func (r *CommandRepository) RunCommand(ctx context.Context, id int) (string, error) {
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
