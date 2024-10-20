package persist

import (
	"context"
	"fmt"

	"github.com/Sreeram-ganesan/my-blog/internal/core/app"
	"github.com/Sreeram-ganesan/my-blog/internal/core/outport"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const createTablesSql =
/*language=postgresql*/ `
CREATE TABLE IF NOT EXISTS contacts(
    id BIGSERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS phones(
    id BIGSERIAL PRIMARY KEY,
    type TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    contact_id BIGINT NOT NULL REFERENCES contacts(id)
);
CREATE TABLE IF NOT EXISTS blogs(
	id BIGSERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	author TEXT NOT NULL,
	content TEXT NOT NULL
);
`

type dbAdapter struct {
	db *sqlx.DB
}

// NewPersistence connects to PostgreSQL database and returns Persistence interface that wraps database reference
func NewPersistence(cfg *app.Config) outport.Persistence {
	dbcfg := cfg.Database
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s connect_timeout=%d",
		dbcfg.Host, dbcfg.Port, dbcfg.Name, dbcfg.User, dbcfg.Password, dbcfg.Sslmode, dbcfg.ConnectTimeout)
	zap.S().Infoln("establishing connection to database...")
	db, err := sqlx.ConnectContext(context.Background(), "postgres", connStr)
	if err != nil {
		zap.S().Fatalln("error connecting to database:", err)
	}
	zap.S().Infoln("connection to database was successfully established, performing initialization...")

	tx := db.MustBegin()
	defer tx.Rollback()
	tx.MustExec(createTablesSql)
	err = tx.Commit()
	if err != nil {
		zap.S().Fatalln("failed to commit transaction while creating database:", err)
	}
	zap.S().Infoln("db initialization was successfully performed")

	return &dbAdapter{db: db}
}

func (d dbAdapter) DB() *sqlx.DB {
	return d.db
}

func (d dbAdapter) Close() {
	err := d.DB().Close()
	if err != nil {
		fmt.Println("failed to close postgresql database:", err)
	}
}
