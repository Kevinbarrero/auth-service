package db

import (
	"database/sql"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if tx != nil {
// 		return err
// 	}
// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rErr := tx.Rollback(); rErr != nil {
// 			return fmt.Errorf("tx err: %v, rollback err: %v", err, rErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }
