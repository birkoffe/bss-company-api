package api

import (
	"context"
	"fmt"

	"github.com/ozonmp/bss-company-api/internal/model"
	"github.com/ozonmp/bss-company-api/internal/repo"
)

type testRepo struct{}

var Companies = []model.Company{
	{
		ID:      0,
		Name:    "Company_0",
		Address: "Unknown_0",
	},
	{
		ID:      1,
		Name:    "Company_1",
		Address: "Unknown_1",
	},
}

func newTestRepo() repo.Repo {
	return &testRepo{}
}

func (r *testRepo) DescribeCompany(ctx context.Context, companyId uint64) (*model.Company, error) {
	if companyId < 0 || int(companyId) >= len(Companies) {
		return nil, fmt.Errorf("Invalid index %d, max elements of companies - %d.", companyId, len(Companies)-1)
	}

	return &Companies[companyId], nil
}
