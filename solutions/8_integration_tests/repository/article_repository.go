package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/greeflas/itea_golang/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ArticleRepository struct {
	conn *pgxpool.Pool
}

func NewArticleRepository(conn *pgxpool.Pool) *ArticleRepository {
	return &ArticleRepository{conn: conn}
}

func (r *ArticleRepository) Create(ctx context.Context, a *model.Article) error {
	sql := `
INSERT INTO articles (id, title, body, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
`

	_, err := r.conn.Exec(
		ctx,
		sql,
		a.Id,
		a.Title,
		a.Body,
		a.CreatedAt,
		a.UpdatedAt,
	)

	return err
}

func (r *ArticleRepository) FindById(ctx context.Context, id uuid.UUID) (*model.Article, error) {
	sql := `SELECT id, title, body, created_at, updated_at FROM articles WHERE id = $1`
	row := r.conn.QueryRow(ctx, sql, id)

	article := new(model.Article)
	err := row.Scan(
		&article.Id,
		&article.Title,
		&article.Body,
		&article.CreatedAt,
		&article.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *ArticleRepository) Update(ctx context.Context, a *model.Article) error {
	sql := `
UPDATE articles SET (title, body, updated_at) = ($2, $3, $4)
WHERE id = $1
`

	_, err := r.conn.Exec(
		ctx,
		sql,
		a.Id,
		a.Title,
		a.Body,
		a.UpdatedAt,
	)

	return err
}
