package service

import (
	"context"
	"github.com/Bazhenator/bashExecAPI/internal/domain"
)

func (s *Service) ListCommands(ctx context.Context) ([]domain.Command, error) {
	return s.commandRepo.ListCommands(ctx)
}

func (s *Service) GetCommand(ctx context.Context, id int) (*domain.Command, error) {
	return s.commandRepo.GetCommand(ctx, id)
}

func (s *Service) CreateCommand(ctx context.Context, command string) (string, string, error) {
	return s.commandRepo.CreateCommand(ctx, command)
}

func (s *Service) RunCommand(ctx context.Context, id int) (string, error) {
	return s.commandRepo.RunCommand(ctx, id)
}
