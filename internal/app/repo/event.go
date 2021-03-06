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

func (r *repo) Add(event []model.CompanyEvent) error {
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

func (r *repo) Lock(n uint64) ([]model.CompanyEvent, error) {
	ctx := context.Background()

	query := sq.Select("id, company_id, type, status, payload, updated").From("company_events").
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{"status": model.Deferred}).Limit(n)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var events []model.CompanyEvent
	err = r.db.SelectContext(ctx, &events, s, args...)
	if err != nil {
		return nil, err
	}

	eventIDs := make([]uint64, 0, len())
	for _, v := range events {
		eventIDs = append(eventIDs, v.ID)
	}

	upadateQuery := sq.Update("company_events").PlaceholderFormat(sq.Dollar).
		Set("status", model.Processed).Where(sq.Eq{"id": eventIDs})

	s, args, err = upadateQuery.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *repo) Unlock(eventIDs []uint64) error {
	ctx := context.Background()

	query := sq.Update("company_events").PlaceholderFormat(sq.Dollar).
		Set("status", model.Deferred).Where(sq.Eq{"id": eventIDs})

	s, args, err := query.ToSql()
	if err != nil {
		return nil
	}

	_, err = r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return err
	}

	return nil
}
