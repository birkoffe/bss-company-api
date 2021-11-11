package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/bss-company-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.CompanyEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.CompanyEvent) error
	Remove(eventIDs []uint64) error
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) repo {
	return &repo{db: db}
}

func (r *repo) Add(event []model.Company) error {
	return nil
}

func (r *repo) Remove(eventIDs []uint64) error {
	ctx := context.Background()

	query := sq.Delete("company_events").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": eventIDs})

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return err
	}

	return nil
}
