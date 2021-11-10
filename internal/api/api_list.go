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

	companies, err := o.repo.ListCompany(ctx, 0, 10)
	if err != nil {
		log.Error().Err(err).Msg("ListCompanyV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if companies == nil {
		log.Debug().Msg("No companies")
		return nil, status.Error(codes.NotFound, "Companies not found")
	}

	log.Debug().Msg("ListCompanyV1 - success")

	response := &pb.ListCompanyV1Response{}

	for _, comp := range companies {
		response.Company = append(response.Company, &pb.Company{Id: comp.ID, Name: comp.Name, Address: comp.Address})
	}

	return response, nil
}
