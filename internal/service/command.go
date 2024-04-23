package service

import (
	"SQLbash/internal/domain"
	"context"
	"fmt"
	"os/exec"
)

func (s *Service) GetCommands(ctx context.Context) ([]domain.Command, error) {
	return s.commandRepo.ListCommands(ctx)
}

func (s *Service) AddCommand(ctx context.Context, command string) (int, error) {
	err := s.commandRepo.AddCommand(ctx, command)
	if err != nil {
		return 0, fmt.Errorf("failed to add command: %w", err)
	}
	// Implement logic to return the ID of the newly added command
	return 0, nil
}

func (s *Service) RunCommand(ctx context.Context, command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %w", err)
	}
	return string(output), nil
}
