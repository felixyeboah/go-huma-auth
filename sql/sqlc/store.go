package db

import (
	"context"
	"database/sql"
)

func ExecTX(ctx context.Context, db *sql.DB, fn func(tx *sql.Tx) error) error {
	if db == nil {
		return nil // Return nil if dependencies are not provided
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	return tx.Commit()
}
