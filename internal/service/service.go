package service

import "bashExecAPI/internal/repository"

type Service struct {
	commandRepo repository.ICommandRepository
}

func NewService(repositories *repository.CommandRepository) *Service {
	return &Service{
		commandRepo: repositories.CommandRepository,
	}
}
