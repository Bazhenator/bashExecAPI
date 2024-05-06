package service

import "github.com/Bazhenator/bashExecAPI/internal/repository"

type Services struct {
	commandRepo repository.ICommandRepository
	dbRepo      repository.IDataBaseRepository
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		commandRepo: repositories.CommandRepository,
		dbRepo:      repositories.DBRepository,
	}
}
