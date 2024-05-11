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
		// Добавьте другие тестовые случаи здесь, если необходимо
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			db, mock, err := sqlxmock.Newx()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			// Ожидание запроса DELETE FROM commands WHERE id = $1
			mock.ExpectExec(`^DELETE FROM commands WHERE id = \$1$`).WithArgs(tc.id).WillReturnResult(sqlxmock.NewResult(0, 1))

			// Если rowCount равно 0, то ожидание запроса SELECT COUNT(*) FROM commands вернет 0
			// В противном случае, будет возвращено значение rowCount
			mock.ExpectQuery(`SELECT COUNT\(\*\) FROM commands`).WillReturnRows(sqlxmock.NewRows([]string{"count"}).AddRow(tc.rowCount))

			if tc.expectResetSeq {
				// Ожидание запроса ALTER SEQUENCE commands_id_seq RESTART WITH 1
				mock.ExpectExec(`^ALTER SEQUENCE commands_id_seq RESTART WITH 1$`).WillReturnResult(sqlxmock.NewResult(0, 1))
			}

			repo := NewDBRepository(&provider.Provider{DB: db})

			// Вызов метода DeleteRow
			err = repo.DeleteRow(context.Background(), tc.id)

			// Проверка ошибки
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error: '%v', got: '%v'", tc.expectedError, err)
			}

			// Проверка, что все ожидания выполнены
			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
