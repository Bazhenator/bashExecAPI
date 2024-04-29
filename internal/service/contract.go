package service

import (
	"bashExecAPI/internal/domain"
	"context"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_service.go . ICommand

type ICommand interface {
	GetCommands(ctx context.Context) ([]domain.Command, error)
	GetCommand(ctx context.Context, id int) (*domain.Command, error)
	CreateCommand(ctx context.Context, command string) (string, error)
	RunCommand(ctx context.Context, id string) (string, error)
	StopCommand(ctx context.Context, id string) error
}
