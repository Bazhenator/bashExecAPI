package repository

import (
	provider "github.com/Bazhenator/bashExecAPI/internal/db"
	db "github.com/Bazhenator/bashExecAPI/internal/repository/psql"
)

type Repositories struct {
	CommandRepository ICommandRepository
	DBRepository      IDataBaseRepository
}

func NewRepositories(provider *provider.Provider) *Repositories {
	return &Repositories{
		CommandRepository: db.NewCommandRepository(provider),
		DBRepository:      db.NewDBRepository(provider),
	}
}
