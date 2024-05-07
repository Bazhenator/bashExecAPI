package service

import (
	"context"
	"github.com/Bazhenator/bashExecAPI/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_service.go . ICommand,IDataBase

type ICommand interface {
	CreateCommand(ctx context.Context, command string) (string, error)
	ListCommands(ctx context.Context) ([]domain.Command, error)
	GetCommand(ctx context.Context, id int) (*domain.Command, error)
	RunCommand(ctx context.Context, id int) (string, error)
}

type IDataBase interface {
	DeleteAllRows(ctx context.Context) error
	DeleteRow(ctx context.Context, id int) error
}
