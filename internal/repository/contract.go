package repository

import (
	"context"

	"bashExecAPI/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_repository.go . ICommandRepository

type ICommandRepository interface {
	ListCommands(ctx context.Context) ([]domain.Command, error)
	GetCommand(ctx context.Context, id int) (*domain.Command, error)
	CreateCommand(ctx context.Context, command string) (string, string, error)
	RunCommand(ctx context.Context, id int) (string, error)
	StopCommand(ctx context.Context, id int) error
}
