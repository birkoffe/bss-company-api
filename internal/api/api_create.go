package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/bss-company-api/pkg/bss-company-api"
)

func (o *CompanyAPI) CreateCompanyV1(
	ctx context.Context,
	req *pb.CreateCompanyV1Request,
) (*pb.CreateCompanyV1Response, error) {
	log.Debug().Msg("CreateCompanyV1")

	return nil, status.Error(codes.Internal, "not implemented")
}
