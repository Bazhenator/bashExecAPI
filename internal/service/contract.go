package service

import (
	"SQLbash/internal/domain"
	"context"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_service.go . ICommand

type ICommand interface {
	ListCommands(ctx context.Context) ([]domain.Command, error)
	GetCommand(ctx context.Context, id int) (*domain.Command, error)
	AddCommand(ctx context.Context, command string) error
}
