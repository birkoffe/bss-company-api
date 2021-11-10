package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/ozonmp/bss-company-api/pkg/bss-company-api"
)

func (o *CompanyAPI) ListCompanyV1(
	ctx context.Context,
	req *emptypb.Empty,
) (*pb.ListCompanyV1Response, error) {
	log.Debug().Msg("ListCompanyV1")

	return nil, status.Error(codes.Internal, "not implemented")
}
