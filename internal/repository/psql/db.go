package psql

import (
	"context"
	"fmt"
	provider "github.com/Bazhenator/bashExecAPI/internal/db"
	log "github.com/sirupsen/logrus"
)

type DataBaseRepository struct {
	db *provider.Provider
}

func NewDBRepository(provider *provider.Provider) *DataBaseRepository {
	return &DataBaseRepository{
		db: provider,
	}
}

func (r *DataBaseRepository) DeleteAllRows(ctx context.Context) error {
	query := `TRUNCATE TABLE commands`
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		log.Error(fmt.Errorf("failed to delete all rows: %w", err))
		return err
	}

	query = `ALTER SEQUENCE commands_id_seq RESTART WITH 1`
	_, err = r.db.ExecContext(ctx, query)
	if err != nil {
		log.Error(fmt.Errorf("failed to reset sequence: %w", err))
		return err
	}

	return nil
}

func (r *DataBaseRepository) DeleteRow(ctx context.Context, id int) error {
	query := `DELETE FROM commands WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Error(fmt.Errorf("failed to delete row with id %d: %w", id, err))
		return err
	}

	if id == 1 {
		var count int
		err = r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM commands").Scan(&count)
		if err != nil {
			log.Error(fmt.Errorf("failed to count rows: %w", err))
			return err
		}

		if count == 0 {
			query = `ALTER SEQUENCE commands_id_seq RESTART WITH 1`
			_, err = r.db.ExecContext(ctx, query)
			if err != nil {
				log.Error(fmt.Errorf("failed to reset sequence: %w", err))
				return err
			}
		}
	}

	return nil
}
