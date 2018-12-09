package db

import (
	"context"
	"database/sql"
	"memo_sample/infra/database"
)

var dbm *database.DBM

func init() {
	dbm = database.GetDBM()
}

// begin begin transaction
func begin(ctx context.Context) (context.Context, error) {
	return (*dbm).Begin(ctx)
}

// rollback rollback transaction
func rollback(ctx context.Context) (context.Context, error) {
	return (*dbm).Rollback(ctx)
}

// commit commit transaction
func commit(ctx context.Context) (context.Context, error) {
	return (*dbm).Commit(ctx)
}

// prepare prepare statement
func prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	return (*dbm).Prepare(ctx, query)
}
