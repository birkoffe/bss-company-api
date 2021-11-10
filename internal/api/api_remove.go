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

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveCompanyV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := o.repo.RemoveCompany(ctx, req.CompanyId)
	if err != nil {
		log.Error().Err(err).Msg("RemoveCompanyV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("RemoveCompanyV1 - success")

	return &pb.RemoveCompanyV1Response{
		Result: result,
	}, nil
}
