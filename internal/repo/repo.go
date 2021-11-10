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
	RemoveCompany(ctx context.Context, companyID uint64) (bool, error)
	ListCompany(ctx context.Context, offset uint64, count uint64) ([]model.Company, error)
	AddCompany(ctx context.Context, company *model.Company) (*model.Company, error)
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
		Where(sq.And{sq.Eq{"id": companyID}, sq.Eq{"removed": false}})

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

func (r *repo) RemoveCompany(ctx context.Context, companyID uint64) (bool, error) {
	query := sq.Update("company").Set("removed", true).Set("updated", "now()").
		Where(sq.Eq{"id": companyID}).PlaceholderFormat(sq.Dollar)

	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	_, err = r.db.QueryContext(ctx, s, args...)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repo) ListCompany(ctx context.Context, offset uint64, count uint64) ([]model.Company, error) {
	query := sq.Select("id, name, address").From("company").Where(sq.Eq{"removed": false}).Limit(count).
		PlaceholderFormat(sq.Dollar)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var companies []model.Company
	err = r.db.SelectContext(ctx, &companies, s, args...)
	if err != nil {
		return nil, err
	}

	if len(companies) == 0 {
		return nil, nil
	}

	return companies, err
}

func (r *repo) AddCompany(ctx context.Context, company *model.Company) (*model.Company, error) {
	return nil, nil
}
