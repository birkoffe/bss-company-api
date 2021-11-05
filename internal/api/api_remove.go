package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/bss-company-api/pkg/bss-company-api"
)

func (o *CompanyAPI) RemoveCompanyV1(
	ctx context.Context,
	req *pb.RemoveCompanyV1Request,
) (*pb.RemoveCompanyV1Response, error) {
	log.Debug().Msg("RemoveCompanyV1")

	return nil, status.Error(codes.Internal, "not implemented")
}
