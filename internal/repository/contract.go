package repository

import (
	"context"

	"bashExecAPI/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_repository.go . ICommandRepository

type ICommandRepository interface {
	GetCommands(ctx context.Context) ([]domain.Command, error)
	GetCommand(ctx context.Context, id string) (*domain.Command, error)
	CreateCommand(ctx context.Context, command string) (string, error)
	StopCommand(ctx context.Context, id string) error
}
