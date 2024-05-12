//go:build unit
// +build unit

package psql

import (
	"context"
	"errors"
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
	testCases := []struct {
		description    string
		id             int
		rowCount       int
		expectResetSeq bool
		expectedError  error
	}{
		{
			description:    "Delete row with id 1 when table is empty",
			id:             1,
			rowCount:       0,
			expectResetSeq: true,
			expectedError:  nil,
		},
		{
			description:    "Delete row with id 1 when table has more rows",
			id:             1,
			rowCount:       2,
			expectResetSeq: false,
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			db, mock, err := sqlxmock.Newx()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectExec(`^DELETE FROM commands WHERE id = \$1$`).WithArgs(tc.id).WillReturnResult(sqlxmock.NewResult(0, 1))

			mock.ExpectQuery(`SELECT COUNT\(\*\) FROM commands`).WillReturnRows(sqlxmock.NewRows([]string{"count"}).AddRow(tc.rowCount))

			if tc.expectResetSeq {
				mock.ExpectExec(`^ALTER SEQUENCE commands_id_seq RESTART WITH 1$`).WillReturnResult(sqlxmock.NewResult(0, 1))
			}

			repo := NewDBRepository(&provider.Provider{DB: db})

			err = repo.DeleteRow(context.Background(), tc.id)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error: '%v', got: '%v'", tc.expectedError, err)
			}

			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
