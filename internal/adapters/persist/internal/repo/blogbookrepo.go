package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type BlogBookRepo struct {
	db                 *sqlx.DB
	insertBlogStmt     *sqlx.NamedStmt
	selectBlogByIdStmt *sqlx.NamedStmt
}

type BlogEntity struct {
	ID      int64
	Title   string
	Content string
	Author  string
}

func NewBlogBookRepo(db *sqlx.DB) *BlogBookRepo {
	return &BlogBookRepo{
		db:                 db,
		insertBlogStmt:     MustPrepareNamed(db, insertBlogSql),
		selectBlogByIdStmt: MustPrepareNamed(db, selectBlogByIdSql),
	}
}

func (r *BlogBookRepo) AddBlog(ctx context.Context, blog *BlogEntity) (*BlogEntity, error) {
	tx := r.db.MustBeginTx(ctx, nil)
	defer tx.Rollback()

	var err error
	newBlog := *blog
	newBlog.ID, err = ExecNamedStmtReturningLastInsertId(ctx, tx.NamedStmtContext(ctx, r.insertBlogStmt), map[string]any{
		"title":   blog.Title,
		"author":  blog.Author,
		"content": blog.Content,
	})
	if err != nil {
		err = fmt.Errorf("error inserting blog into database: %w", err)
		zap.S().Errorln(err)
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		err = fmt.Errorf("error committing transaction: %w", err)
		zap.S().Errorln(err)
		return nil, err
	}
	return &newBlog, nil
}
