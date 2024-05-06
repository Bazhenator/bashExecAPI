package repository

import (
	provider "github.com/Bazhenator/bashExecAPI/internal/db"
	db "github.com/Bazhenator/bashExecAPI/internal/repository/psql"
)

type CommandRepository struct {
	CommandRepository ICommandRepository
}

func NewCommandRepository(provider *provider.Provider) *CommandRepository {
	return &CommandRepository{
		CommandRepository: db.NewCommandRepository(provider),
	}
}
