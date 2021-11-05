package api

import (
	"context"
	"testing"

	bss_company_api "github.com/ozonmp/bss-company-api/pkg/bss-company-api"
	"github.com/stretchr/testify/assert"
)

func TestDescribeValid(t *testing.T) {
	CompanyAPI := NewCompanyAPI(newTestRepo())

	t.Parallel()

	response, err := CompanyAPI.DescribeCompanyV1(
		context.Background(),
		&bss_company_api.DescribeCompanyV1Request{
			CompanyId: 1,
		},
	)

	responseId := response.GetValue().GetId()
	assert.Equal(t, uint64(1), responseId, "Not excepted response")
	assert.Equal(t, nil, err, "Err must be nil")
}

func TestDescribeLessOne(t *testing.T) {
	CompanyAPI := NewCompanyAPI(newTestRepo())

	t.Parallel()

	_, err := CompanyAPI.DescribeCompanyV1(
		context.Background(),
		&bss_company_api.DescribeCompanyV1Request{
			CompanyId: 0,
		},
	)

	errEx := "rpc error: code = InvalidArgument desc = invalid DescribeCompanyV1Request.CompanyId: value must be greater than 0"
	errDesc := err.Error()
	assert.Equal(t, errEx, errDesc, "")
}

func TestDescribeNoCompany(t *testing.T) {
	CompanyAPI := NewCompanyAPI(newTestRepo())

	t.Parallel()

	_, err := CompanyAPI.DescribeCompanyV1(
		context.Background(),
		&bss_company_api.DescribeCompanyV1Request{
			CompanyId: 2,
		},
	)

	errEx := "rpc error: code = Internal desc = No company with idx: 2"
	errDesc := err.Error()
	assert.Equal(t, errEx, errDesc, "")
}
