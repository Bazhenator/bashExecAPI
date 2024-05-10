//go:build unit
// +build unit

package psql

import (
	"context"
	"fmt"
	provider "github.com/Bazhenator/bashExecAPI/internal/db"
	"github.com/Bazhenator/bashExecAPI/internal/domain"
	repos "github.com/Bazhenator/bashExecAPI/internal/repository"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"strconv"
	"testing"
)

func TestCommandRepository_CreateCommand(t *testing.T) {
	mockCommand := "echo hello world"
	mockResult := "hello world\n"
	mockID := 1

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`^INSERT INTO commands \(command\) VALUES \(\$1\) RETURNING id$`).
		WithArgs(mockCommand).
		WillReturnRows(sqlxmock.NewRows([]string{"id"}).AddRow(mockID))

	mock.ExpectQuery(`^SELECT \* FROM commands WHERE id = \? LIMIT 1$`).
		WithArgs(mockID).
		WillReturnRows(sqlxmock.NewRows([]string{"id", "command"}).AddRow(mockID, mockCommand))

	mock.ExpectExec(`^UPDATE commands SET result = \$1 WHERE id = \$2$`).
		WithArgs(mockResult, mockID).
		WillReturnResult(sqlxmock.NewResult(1, 1))

	repo := repos.NewRepositories(&provider.Provider{DB: db})
	result, id, err := repo.CommandRepository.CreateCommand(context.Background(), mockCommand)

	err = repo.DBRepository.DeleteAllRows(context.Background())
	if err != nil {
		log.Error(fmt.Errorf("failed to delete rows: %v", err))
	}

	assert.NoError(t, err)
	assert.Equal(t, mockResult, result)
	assert.Equal(t, strconv.Itoa(mockID), id)
}

func TestCommandRepository_ListCommands(t *testing.T) {
	commandRows := sqlxmock.NewRows([]string{"id", "command"}).
		AddRow(1, "echo test 1").
		AddRow(2, "echo test 2")

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`^SELECT \* FROM commands$`).
		WillReturnRows(commandRows)

	repo := repos.NewRepositories(&provider.Provider{DB: db})
	commands, err := repo.CommandRepository.ListCommands(context.Background())

	err = repo.DBRepository.DeleteAllRows(context.Background())
	if err != nil {
		log.Error(fmt.Errorf("failed to delete rows: %v", err))
	}

	assert.NoError(t, err)
	assert.Len(t, commands, 2)
}

func TestCommandRepository_GetCommand(t *testing.T) {
	commandRows := sqlxmock.NewRows([]string{"id", "command"}).
		AddRow(1, "echo test")

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`^SELECT \* FROM commands WHERE id = \? LIMIT 1$`).
		WithArgs(1).
		WillReturnRows(commandRows)

	repo := repos.NewRepositories(&provider.Provider{DB: db})
	command, err := repo.CommandRepository.GetCommand(context.Background(), 1)

	err = repo.DBRepository.DeleteAllRows(context.Background())
	if err != nil {
		log.Error(fmt.Errorf("failed to delete rows: %v", err))
	}

	assert.NoError(t, err)
	assert.NotNil(t, command)
}

func TestCommandRepository_RunCommand(t *testing.T) {
	mockCommand := &domain.Command{
		ID:      1,
		Command: "echo hello world",
	}

	output := "hello world\n"
	outputRows := sqlxmock.NewRows([]string{"output"}).
		AddRow(output)

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`^SELECT \* FROM commands WHERE id = \? LIMIT 1$`).
		WithArgs(1).
		WillReturnRows(outputRows)

	repo := repos.NewRepositories(&provider.Provider{DB: db})
	actualOutput, err := repo.CommandRepository.RunCommand(context.Background(), mockCommand.ID)

	err = repo.DBRepository.DeleteAllRows(context.Background())
	if err != nil {
		log.Error(fmt.Errorf("failed to delete rows: %v", err))
	}

	assert.NoError(t, err)
	assert.Equal(t, output, actualOutput)
}
