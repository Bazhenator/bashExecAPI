package service

import (
	"bashExecAPI/internal/domain"
	"context"
)

func (s *Service) GetCommands(ctx context.Context) ([]domain.Command, error) {

	return s.commandRepo.GetCommands(ctx)
}

func (s *Service) GetCommand(ctx context.Context, id string) (*domain.Command, error) {
	return s.commandRepo.GetCommand(ctx, id)
}

func (s *Service) CreateCommand(ctx context.Context, command string) (string, error) {
	return s.commandRepo.CreateCommand(ctx, command)
}

func (s *Service) RunCommand(ctx context.Context, id string) (string, error) {
	return s.commandRepo.RunCommand(ctx, id)
}

func (s *Service) StopCommand(ctx context.Context, id string) error {
	return s.commandRepo.StopCommand(ctx, id)
}
