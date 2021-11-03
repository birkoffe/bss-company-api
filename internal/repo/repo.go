package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/bss-company-api/internal/model"
)

// Repo is DAO for company
type Repo interface {
	Describecompany(ctx context.Context, companyID uint64) (*model.company, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) Describecompany(ctx context.Context, companyID uint64) (*model.company, error) {
	return nil, nil
}
