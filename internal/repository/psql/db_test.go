package psql

import (
	"context"
	provider "github.com/Bazhenator/bashExecAPI/internal/db"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestDataBaseRepository_DeleteAllRows(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBRepository(&provider.Provider{DB: db})

	mock.ExpectExec(`^TRUNCATE TABLE commands$`).WillReturnResult(sqlxmock.NewResult(0, 0))
	mock.ExpectExec(`^ALTER SEQUENCE commands_id_seq RESTART WITH 1$`).WillReturnResult(sqlxmock.NewResult(0, 0))

	err = repo.DeleteAllRows(context.Background())

	assert.NoError(t, err)
}

func TestDataBaseRepository_DeleteRow(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBRepository(&provider.Provider{DB: db})
	mock.ExpectExec(`^DELETE FROM commands WHERE id = \$1$`).WillReturnResult(sqlxmock.NewResult(0, 0))

	err = repo.DeleteRow(context.Background(), 1)

	assert.NoError(t, err)
}
