package repository

import (
	provider "SQLbash/internal/db"
	db "SQLbash/internal/repository/psql"
)

type CommandRepository struct {
	CommandRepository ICommandRepository
}

func NewCommandRepository(provider *provider.Provider) *CommandRepository {
	return &CommandRepository{
		CommandRepository: db.NewCommandRepository(provider),
	}
}
