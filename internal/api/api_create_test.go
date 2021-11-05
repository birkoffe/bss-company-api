package api

import (
	"context"
	"testing"

	bss_company_api "github.com/ozonmp/bss-company-api/pkg/bss-company-api"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	CompanyAPI := NewCompanyAPI(newTestRepo())

	t.Parallel()

	_, err := CompanyAPI.CreateCompanyV1(
		context.Background(),
		&bss_company_api.CreateCompanyV1Request{
			CompanyName: "Horizpn Ltd",
			AddressName: "Nowhere",
		},
	)

	errEx := "rpc error: code = Internal desc = not implemented"
	errDesc := err.Error()
	assert.Equal(t, errEx, errDesc, "")
}
