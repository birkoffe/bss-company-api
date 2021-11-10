package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/bss-company-api/internal/model"
)

// Repo is DAO for company
type Repo interface {
	DescribeCompany(ctx context.Context, companyID uint64) (*model.Company, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeCompany(ctx context.Context, companyID uint64) (*model.Company, error) {
	query := sq.Select("id, name, address, removed, created, updated").
		PlaceholderFormat(sq.Dollar).
		From("company").
		Where(sq.Eq{"id": companyID})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var company []model.Company
	err = r.db.SelectContext(ctx, &company, s, args...)
	if err != nil {
		return nil, err
	}

	if len(company) == 0 {
		return nil, nil
	}

	return &company[0], err
}
