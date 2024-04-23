package repository

import (
	"context"

	"SQLbash/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_repository.go . ICommandRepository

type ICommandRepository interface {
	ListCommands(ctx context.Context) ([]domain.Command, error)
	GetCommand(ctx context.Context, id int) (*domain.Command, error)
	AddCommand(ctx context.Context, command string) error
}
