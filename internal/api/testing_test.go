package api

import (
	"context"
	"fmt"

	"github.com/ozonmp/bss-company-api/internal/model"
	"github.com/ozonmp/bss-company-api/internal/repo"
)

type testRepo struct{}

var companies = map[uint64]model.Company{
	1: {
		ID:      1,
		Name:    "Company_1",
		Address: "Unknown_1",
	},
}

func newTestRepo() repo.Repo {
	return &testRepo{}
}

func (r *testRepo) DescribeCompany(ctx context.Context, companyId uint64) (*model.Company, error) {
	company, ok := companies[companyId]
	if !ok {
		return nil, fmt.Errorf("No company with idx: %d", companyId)
	}

	return &company, nil
}
