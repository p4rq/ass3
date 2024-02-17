package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	db      *pgxpool.Pool
	genSQL  squirrel.StatementBuilderType
	options Options
}

type Options struct{}

func New(db *pgxpool.Pool, o Options) *Repository {
	var r = &Repository{
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db:     db,
	}

	r.SetOptions(o)
	return r
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}
