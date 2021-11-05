package api

import (
	"context"

	pb "github.com/ozonmp/bss-company-api/pkg/bss-company-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *CompanyAPI) DescribeCompanyV1(
	ctx context.Context,
	req *pb.DescribeCompanyV1Request,
) (*pb.DescribeCompanyV1Response, error) {
	log.Debug().Msg("DescribeCompanyV1")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeCompanyV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	company, err := o.repo.DescribeCompany(ctx, req.CompanyId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeCompanyV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if company == nil {
		log.Debug().Uint64("companyId", req.CompanyId).Msg("company not found")
		totalCompanyNotFound.Inc()

		return nil, status.Error(codes.NotFound, "company not found")
	}

	log.Debug().Msg("DescribeCompanyV1 - success")

	return &pb.DescribeCompanyV1Response{
		Value: &pb.Company{
			Id:      company.ID,
			Name:    company.Name,
			Address: company.Address,
		},
	}, nil
}
