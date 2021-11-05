package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestList(t *testing.T) {
	CompanyAPI := NewCompanyAPI(newTestRepo())

	t.Parallel()

	_, err := CompanyAPI.ListCompanyV1(
		context.Background(),
		&emptypb.Empty{},
	)

	err_ex := "rpc error: code = Internal desc = not implemented"
	err_desc := err.Error()
	assert.Equal(t, err_ex, err_desc, "")
}
