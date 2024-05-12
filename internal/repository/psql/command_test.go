//go:build unit
// +build unit

package psql

import (
	"context"
	provider "github.com/Bazhenator/bashExecAPI/internal/db"
	"github.com/Bazhenator/bashExecAPI/internal/domain"
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

	mock.ExpectQuery(`^INSERT INTO commands \(command\) VALUES \(\$1\) RETURNING id;$`).
		WithArgs(mockCommand).
		WillReturnRows(sqlxmock.NewRows([]string{"id"}).AddRow(mockID))

	mock.ExpectQuery(`^SELECT id, command, result FROM commands WHERE id = (.+)$`).
		WithArgs(mockID).
		WillReturnRows(sqlxmock.NewRows([]string{"id", "command", "result"}).AddRow(mockID, mockCommand, mockResult))

	mock.ExpectExec(`^UPDATE commands SET result = \$1 WHERE id = \$2;$`).
		WithArgs(mockResult, mockID).
		WillReturnResult(sqlxmock.NewResult(1, 1))

	repoC := NewCommandRepository(&provider.Provider{DB: db})
	result, id, err := repoC.CreateCommand(context.Background(), mockCommand)

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

	repoC := NewCommandRepository(&provider.Provider{DB: db})
	commands, err := repoC.ListCommands(context.Background())

	assert.NoError(t, err)
	assert.Len(t, commands, 2)
}

func TestCommandRepository_GetCommand(t *testing.T) {
	var commandRow = sqlxmock.NewRows([]string{"id", "command", "result"}).
		AddRow(1, "echo Hello World!", "Hello World\n")

	output := "Hello World\n"

	var expected = domain.Command{
		1,
		"echo Hello World!",
		&output,
	}

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	t.Cleanup(func() {
		db.Close()
	})

	mock.ExpectQuery(`^SELECT id, command, result FROM commands WHERE id = (.+)$`).
		WithArgs(1).
		WillReturnRows(commandRow)

	repo := NewCommandRepository(&provider.Provider{
		DB: db,
	})

	command, err := repo.GetCommand(context.Background(), 1)
	if assert.NoError(t, err) {
		assert.Equal(t, command, &expected)
	}
}

func TestCommandRepository_RunCommand(t *testing.T) {
	output := "hello world\n"

	mockCommand := &domain.Command{
		ID:      1,
		Command: "echo hello world",
		Result:  &output,
	}

	var commandRow = sqlxmock.NewRows([]string{"id", "command", "result"}).
		AddRow(1, "echo hello world", "hello world\n")

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`^SELECT id, command, result FROM commands WHERE id = (.+)$`).
		WithArgs(1).
		WillReturnRows(commandRow)

	repoC := NewCommandRepository(&provider.Provider{DB: db})
	actualOutput, err := repoC.RunCommand(context.Background(), mockCommand.ID)

	if assert.NoError(t, err) {
		assert.Equal(t, actualOutput, output)
	}
}
