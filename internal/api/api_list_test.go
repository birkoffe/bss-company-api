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

	errEx := "rpc error: code = Internal desc = not implemented"
	errDesc := err.Error()
	assert.Equal(t, errEx, errDesc, "")
}
