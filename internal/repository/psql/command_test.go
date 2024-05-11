//go:build unit
// +build unit

package psql

import (
	"context"
	provider "github.com/Bazhenator/bashExecAPI/internal/db"
	"github.com/Bazhenator/bashExecAPI/internal/domain"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
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

	repoC := NewCommandRepository(&provider.Provider{DB: db})
	result, id, err := repoC.CreateCommand(context.Background(), mockCommand)

	assert.NoError(t, err)
	/*assert.Equal(t, mockResult, result)
	assert.Equal(t, strconv.Itoa(mockID), id)*/
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

	repoC := NewCommandRepository(&provider.Provider{DB: db})
	command, err := repoC.GetCommand(context.Background(), 1)

	assert.NoError(t, err)
	/*assert.NotNil(t, command)*/
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

	repoC := NewCommandRepository(&provider.Provider{DB: db})
	actualOutput, err := repoC.RunCommand(context.Background(), mockCommand.ID)

	assert.NoError(t, err)
	/*assert.Equal(t, output, actualOutput)*/
}
