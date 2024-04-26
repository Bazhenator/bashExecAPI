package repository

import (
	provider "bashExecAPI/internal/db"
	db "bashExecAPI/internal/repository/psql"
)

type CommandRepository struct {
	CommandRepository ICommandRepository
}

func NewCommandRepository(provider *provider.Provider) *CommandRepository {
	return &CommandRepository{
		CommandRepository: db.NewCommandRepository(provider),
	}
}
