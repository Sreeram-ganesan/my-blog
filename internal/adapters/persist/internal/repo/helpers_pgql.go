package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func ExecNamedStmtReturningLastInsertId(ctx context.Context, stmt *sqlx.NamedStmt, arg any) (int64, error) {
	var id int64
	err := stmt.QueryRowContext(ctx, arg).Scan(&id)
	return id, err
}
