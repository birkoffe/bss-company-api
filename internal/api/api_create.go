package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozonmp/bss-company-api/internal/model"
	pb "github.com/ozonmp/bss-company-api/pkg/bss-company-api"
)

func (o *CompanyAPI) CreateCompanyV1(
	ctx context.Context,
	req *pb.CreateCompanyV1Request,
) (*pb.CreateCompanyV1Response, error) {
	log.Debug().Msg("CreateCompanyV1")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateCompanyV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	company, err := o.repo.AddCompany(ctx, &model.Company{
		Name:    req.GetCompanyName(),
		Address: req.GetAddressName(),
	})
	if err != nil {
		log.Error().Err(err).Msg("CreateCompanyV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateCompanyV1 - success")

	return &pb.CreateCompanyV1Response{
		Company: &pb.Company{
			Id:      company.ID,
			Name:    company.Name,
			Address: company.Address,
		},
	}, nil
}
